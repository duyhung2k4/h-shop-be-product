package controller

import (
	"app/config"
	"app/dto/request"
	"app/grpc/proto"
	"app/model"
	"app/service"
	"app/utils"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type productController struct {
	productService      service.ProductService
	clientFileGRPC      proto.FileServiceClient
	warehouseService    proto.WarehouseServiceClient
	shopService         proto.ShopServiceClient
	queueProductService service.QueueProductService
	jwtUtils            utils.JwtUtils
	redisUtils          utils.RedisUtils
}

type ProductController interface {
	GetProductByProfileId(w http.ResponseWriter, r *http.Request)
	GetProductById(w http.ResponseWriter, r *http.Request)
	CreateProduct(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
	DeleteProduct(w http.ResponseWriter, r *http.Request)

	Heart(w http.ResponseWriter, r *http.Request)
	IsHeart(w http.ResponseWriter, r *http.Request)
	GetHeart(w http.ResponseWriter, r *http.Request)

	Cart(w http.ResponseWriter, r *http.Request)
	IsCart(w http.ResponseWriter, r *http.Request)
	GetCart(w http.ResponseWriter, r *http.Request)
}

func (c *productController) GetProductByProfileId(w http.ResponseWriter, r *http.Request) {
	mapDataRequest, errMapData := c.jwtUtils.GetMapData(r)
	if errMapData != nil {
		internalServerError(w, r, errMapData)
		return
	}

	profileId := uint(mapDataRequest["profile_id"].(float64))
	products, err := c.productService.GetProductbyProfileId(uint64(profileId))
	if err != nil {
		internalServerError(w, r, err)
		return
	}

	res := Response{
		Data:    products,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func (c *productController) GetProductById(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	productId := params.Get("id")

	if productId == "" {
		badRequest(w, r, errors.New("productId empty"))
		return
	}

	productInRedis, errProductInRedis := c.redisUtils.GetData(productId)
	if errProductInRedis == nil {
		res := Response{
			Data:    productInRedis,
			Message: "OK",
			Status:  200,
			Error:   nil,
		}
		render.JSON(w, r, res)
		return
	}

	product, err := c.productService.GetProductById(productId)
	if err != nil {
		internalServerError(w, r, err)
		return
	}

	res := Response{
		Data:    product,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func (c *productController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product request.CreateProductRequest

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		badRequest(w, r, err)
		return
	}

	for keyDefault := range model.MapDefaultFieldProduct {
		checkExit := false
		for key := range product.InfoProduct {
			if key == keyDefault {
				checkExit = true
				break
			}
		}

		if !checkExit {
			badRequest(w, r, errors.New("missing default field"))
			return
		}
	}

	// Create product
	mapDataRequest, errMapData := c.jwtUtils.GetMapData(r)
	if errMapData != nil {
		internalServerError(w, r, errMapData)
		return
	}

	profileId := uint(mapDataRequest["profile_id"].(float64))

	result, errGetShop := c.shopService.GetShopByProfileId(r.Context(), &proto.GetShopByProfileIdReq{ProfileId: uint64(profileId)})
	if errGetShop != nil {
		internalServerError(w, r, errGetShop)
		return
	}

	product.InfoProduct["profileId"] = profileId
	product.InfoProduct["shopId"] = result.ShopId
	newProduct, errProduct := c.productService.CreateProduct(product.InfoProduct)
	if errProduct != nil {
		internalServerError(w, r, errProduct)
		return
	}

	var errHandle error = nil
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		// Create warehouse
		_, err := c.warehouseService.Insert(context.Background(), &proto.InsertWarehouseReq{
			ProductId: newProduct["_id"].(primitive.ObjectID).Hex(),
		})
		if err != nil {
			errHandle = err
			log.Println("Create warehouse: ", err)
			wg.Done()
			return
		}
		wg.Done()
	}()
	go func() {
		// Add avatar
		if product.Avatar != nil {
			_, err := c.clientFileGRPC.InsertAvatarProduct(context.Background(), &proto.InsertAvatarProductReq{
				Data:      product.Avatar.DataBytes,
				Format:    product.Avatar.Format,
				Name:      product.Avatar.Name,
				ProductId: newProduct["_id"].(primitive.ObjectID).Hex(),
			})

			if err != nil {
				errHandle = err
				wg.Done()
				return
			}
		}
		wg.Done()
	}()
	go func() {
		// Handle file
		if len(product.Files) > 0 {
			postFile, errPostFile := c.clientFileGRPC.InsertFile(context.Background())
			if errPostFile != nil {
				errHandle = errPostFile
				wg.Done()
				return
			}
			for _, file := range product.Files {
				postFile.Send(&proto.InsertFileReq{
					Data:      file.DataBytes,
					Format:    file.Format,
					Name:      file.Name,
					ProductId: newProduct["_id"].(primitive.ObjectID).Hex(),
					TypeModel: string(model.PRODUCT),
				})
			}
			_, errResFile := postFile.CloseAndRecv()
			if errResFile != nil {
				errHandle = errPostFile
				wg.Done()
				return
			}
		}
		wg.Done()
	}()
	wg.Wait()

	if errHandle != nil {
		internalServerError(w, r, errHandle)
		return
	}

	newProduct["_id"] = newProduct["_id"].(primitive.ObjectID).Hex()
	c.queueProductService.PushMessInQueueToElasticSearch(newProduct, string(model.PRODUCT_TO_ELASTIC))

	res := Response{
		Data:    newProduct,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func (c *productController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product request.UpdateProductRequest

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		badRequest(w, r, err)
		return
	}

	// Check permission action with product
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")[1]
	productId := product.InfoProduct["_id"].(string)
	isPermission, errPermission := c.productService.CheckPermissionProduct(productId, tokenString)
	if errPermission != nil {
		internalServerError(w, r, errPermission)
		return
	}
	if isPermission == &model.FALSE_VALUE {
		handleError(w, r, errors.New("not permision"), 401)
		return
	}

	// Update product
	newProduct, err := c.productService.UpdateProduct(product.InfoProduct)
	if err != nil {
		internalServerError(w, r, err)
		return
	}

	// Handle file
	var wg sync.WaitGroup
	var errHandleFile = err
	wg.Add(3)

	go func() {
		if len(product.ListFileIdDeletes) > 0 {
			_, err := c.clientFileGRPC.DeleteFile(context.Background(), &proto.DeleteFileReq{
				Ids: product.ListFileIdDeletes,
			})
			errHandleFile = err
		}

		wg.Done()
	}()

	go func() {
		if product.Avatar != nil {
			_, err := c.clientFileGRPC.InsertAvatarProduct(context.Background(), &proto.InsertAvatarProductReq{
				ProductId: productId,
				Name:      product.Avatar.Name,
				Format:    product.Avatar.Format,
				Data:      product.Avatar.DataBytes,
			})

			errHandleFile = err
		}
		wg.Done()
	}()

	go func() {
		if len(product.Files) > 0 {
			postFile, errPostFile := c.clientFileGRPC.InsertFile(context.Background())
			if errPostFile != nil {
				errHandleFile = errPostFile
				wg.Done()
				return
			}
			for _, file := range product.Files {
				postFile.Send(&proto.InsertFileReq{
					Data:      file.DataBytes,
					Format:    file.Format,
					Name:      file.Name,
					ProductId: newProduct["_id"].(string),
					TypeModel: string(model.PRODUCT),
				})
			}
			resFile, errResFile := postFile.CloseAndRecv()
			if errResFile != nil {
				errHandleFile = errResFile
				wg.Done()
				return
			}
			newProduct["fileIds"] = resFile.FileIds
		}
		wg.Done()
	}()
	wg.Wait()

	if errHandleFile != nil {
		internalServerError(w, r, errHandleFile)
		return
	}

	c.queueProductService.PushMessInQueueToElasticSearch(newProduct, string(model.UPDATE_PRODUCT_TO_ELASTIC))
	if err := c.redisUtils.Cache(newProduct["_id"].(string), newProduct); err != nil {
		internalServerError(w, r, err)
		return
	}

	res := Response{
		Data:    newProduct,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func (c *productController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var product request.DeleteProductRequest

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		badRequest(w, r, err)
		return
	}

	// Check permission action with product
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")[1]
	isPermission, errPermission := c.productService.CheckPermissionProduct(product.ProductId, tokenString)
	if errPermission != nil {
		internalServerError(w, r, errPermission)
		return
	}
	if isPermission == &model.FALSE_VALUE {
		handleError(w, r, errors.New("not permision"), 401)
		return
	}

	// Delete files
	listFileIds, errListFileIds := c.clientFileGRPC.GetFileIdsWithProductId(
		context.Background(),
		&proto.GetFileIdsWithProductIdReq{ProductId: product.ProductId},
	)

	if errListFileIds != nil {
		internalServerError(w, r, errListFileIds)
		return
	}
	if _, err := c.clientFileGRPC.
		DeleteFile(context.Background(), &proto.DeleteFileReq{Ids: listFileIds.Ids}); err != nil {
		internalServerError(w, r, err)
		return
	}

	// Delete product
	err := c.productService.DeleteProduct(product.ProductId)
	if err != nil {
		internalServerError(w, r, err)
		return
	}

	c.queueProductService.PushMessInQueueToElasticSearch(product, string(model.DELETE_PRODUCT_TO_ELASTIC))
	if err := c.redisUtils.Delete(product.ProductId); err != nil {
		internalServerError(w, r, err)
		return
	}

	res := Response{
		Data:    nil,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func (c *productController) Heart(w http.ResponseWriter, r *http.Request) {
	var payload request.HeartRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		badRequest(w, r, err)
		return
	}

	mapDataRequest, errMapData := c.jwtUtils.GetMapData(r)
	if errMapData != nil {
		internalServerError(w, r, errMapData)
		return
	}

	profileId := uint(mapDataRequest["profile_id"].(float64))
	err := c.productService.Heart(payload.ProductId, profileId)
	if err != nil {
		internalServerError(w, r, err)
		return
	}

	res := Response{
		Data:    nil,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func (c *productController) IsHeart(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	productId := params.Get("id")
	if productId == "" {
		badRequest(w, r, errors.New("id empty"))
		return
	}

	mapDataRequest, errMapData := c.jwtUtils.GetMapData(r)
	if errMapData != nil {
		internalServerError(w, r, errMapData)
		return
	}

	profileId := uint(mapDataRequest["profile_id"].(float64))
	isHeart := c.productService.IsHeart(productId, profileId)

	res := Response{
		Data:    isHeart,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func (c *productController) GetHeart(w http.ResponseWriter, r *http.Request) {
	mapDataRequest, errMapData := c.jwtUtils.GetMapData(r)
	if errMapData != nil {
		internalServerError(w, r, errMapData)
		return
	}

	profileId := uint(mapDataRequest["profile_id"].(float64))

	data, err := c.productService.GetHeart(profileId)
	if err != nil {
		internalServerError(w, r, err)
		return
	}

	res := Response{
		Data:    data,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func (c *productController) Cart(w http.ResponseWriter, r *http.Request) {
	var payload request.CartRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		badRequest(w, r, err)
		return
	}

	mapDataRequest, errMapData := c.jwtUtils.GetMapData(r)
	if errMapData != nil {
		internalServerError(w, r, errMapData)
		return
	}

	profileId := uint(mapDataRequest["profile_id"].(float64))
	err := c.productService.Cart(payload.ProductId, profileId)
	if err != nil {
		internalServerError(w, r, err)
		return
	}

	res := Response{
		Data:    nil,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func (c *productController) IsCart(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	productId := params.Get("id")
	if productId == "" {
		badRequest(w, r, errors.New("id empty"))
		return
	}

	mapDataRequest, errMapData := c.jwtUtils.GetMapData(r)
	if errMapData != nil {
		internalServerError(w, r, errMapData)
		return
	}

	profileId := uint(mapDataRequest["profile_id"].(float64))
	isCart := c.productService.IsCart(productId, profileId)

	res := Response{
		Data:    isCart,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func (c *productController) GetCart(w http.ResponseWriter, r *http.Request) {
	mapDataRequest, errMapData := c.jwtUtils.GetMapData(r)
	if errMapData != nil {
		internalServerError(w, r, errMapData)
		return
	}

	profileId := uint(mapDataRequest["profile_id"].(float64))

	data, err := c.productService.GetCart(profileId)
	if err != nil {
		internalServerError(w, r, err)
		return
	}

	res := Response{
		Data:    data,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func NewProductController() ProductController {
	return &productController{
		productService:      service.NewProductService(),
		clientFileGRPC:      proto.NewFileServiceClient(config.GetConnFileGrpc()),
		warehouseService:    proto.NewWarehouseServiceClient(config.GetConnWarehouseGrpc()),
		shopService:         proto.NewShopServiceClient(config.GetConnShopGRPC()),
		queueProductService: service.NewQueueProductService(),
		jwtUtils:            utils.NewJwtUtils(),
		redisUtils:          utils.NewUtilsRedis(),
	}
}
