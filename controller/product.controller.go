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
	queueProductService service.QueueProductService
	utils               utils.JwtUtils
}

type ProductController interface {
	CreateProduct(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
	DeleteProduct(w http.ResponseWriter, r *http.Request)
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

	// Check permission action with shop
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")[1]
	shopId := product.InfoProduct["shopId"].(float64)
	isPermission, errIsPermission := c.productService.CheckPermissionShop(uint(shopId), tokenString)
	if errIsPermission != nil {
		internalServerError(w, r, errIsPermission)
		return
	}
	if isPermission == &model.FALSE_VALUE {
		handleError(w, r, errors.New("not permission"), 404)
		return
	}

	// Create product
	mapDataRequest, errMapData := c.utils.JwtDecode(tokenString)
	if errMapData != nil {
		internalServerError(w, r, errMapData)
		return
	}

	product.InfoProduct["profileId"] = uint(mapDataRequest["profile_id"].(float64))
	newProduct, errProduct := c.productService.CreateProduct(product.InfoProduct)
	if errProduct != nil {
		internalServerError(w, r, errProduct)
		return
	}

	var errHandle error = nil
	var wg sync.WaitGroup
	wg.Add(2)

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
			resFile, errResFile := postFile.CloseAndRecv()
			if errResFile != nil {
				errHandle = errPostFile
				wg.Done()
				return
			}
			newProduct["fileIds"] = resFile.FileIds
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
	wg.Add(2)

	go func() {
		if len(product.ListFileIdDeletes) > 0 {
			_, errDeleteFile := c.clientFileGRPC.DeleteFile(context.Background(), &proto.DeleteFileReq{
				Ids: product.ListFileIdDeletes,
			})
			if errDeleteFile != nil {
				internalServerError(w, r, errDeleteFile)
				return
			}
		}

		wg.Done()
	}()

	go func() {
		if len(product.Files) > 0 {
			postFile, errPostFile := c.clientFileGRPC.InsertFile(context.Background())
			if errPostFile != nil {
				internalServerError(w, r, errPostFile)
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
			resFile, errResFile := postFile.CloseAndRecv()
			if errResFile != nil {
				internalServerError(w, r, errResFile)
				return
			}
			newProduct["fileIds"] = resFile.FileIds
		}
		wg.Done()
	}()
	wg.Wait()

	newProduct["_id"] = newProduct["_id"].(primitive.ObjectID).Hex()
	c.queueProductService.PushMessInQueueToElasticSearch(newProduct, string(model.UPDATE_PRODUCT_TO_ELASTIC))

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

	res := Response{
		Data:    nil,
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
		queueProductService: service.NewQueueProductService(),
		utils:               utils.NewJwtUtils(),
	}
}
