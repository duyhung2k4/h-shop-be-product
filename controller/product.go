package controller

import (
	"app/config"
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

	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type productController struct {
	productService service.ProductService
	clientShopGRPC proto.ShopServiceClient
	clientFileGRPC proto.FileServiceClient
	utils          utils.JwtUtils
}

type ProductController interface {
	CreateProduct(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
	DeleteProduct(w http.ResponseWriter, r *http.Request)
}

func (c *productController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product service.CreateProductPayload

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

	// Check permission action with product
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")[1]
	mapDataRequest, errMapData := c.utils.JwtDecode(tokenString)
	if errMapData != nil {
		internalServerError(w, r, errMapData)
		return
	}
	profileId := uint(mapDataRequest["profile_id"].(float64))
	product.InfoProduct["profileId"] = profileId
	shopId := uint(product.InfoProduct["shopId"].(float64))

	resPermissionShop, errPermissionShop := c.clientShopGRPC.CheckShopPermission(
		context.Background(),
		&proto.CheckShopPermissionReq{
			ShopId:    uint64(shopId),
			ProfileId: uint64(profileId),
		},
	)
	if errPermissionShop != nil {
		internalServerError(w, r, errPermissionShop)
		return
	}
	if !resPermissionShop.IsPermission {
		handleError(w, r, errors.New("not permisson"), 401)
		return
	}
	//

	// Create product
	newProduct, errProduct := c.productService.CreateProduct(product.InfoProduct)
	if errProduct != nil {
		internalServerError(w, r, errProduct)
		return
	}

	log.Println(newProduct)

	// Handle file
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

	res := Response{
		Data:    newProduct,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func (c *productController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product service.UpdateProductPayload

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		badRequest(w, r, err)
		return
	}

	// Check permission action with product
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")[1]
	mapDataRequest, errMapData := c.utils.JwtDecode(tokenString)
	if errMapData != nil {
		internalServerError(w, r, errMapData)
		return
	}
	profileId := uint(mapDataRequest["profile_id"].(float64))
	productId := product.InfoProduct["_id"].(string)
	shopId := uint(product.InfoProduct["shopId"].(float64))

	resPermissionShop, errPermissionShop := c.clientShopGRPC.CheckShopPermission(
		context.Background(),
		&proto.CheckShopPermissionReq{
			ShopId:    uint64(shopId),
			ProfileId: uint64(profileId),
		},
	)
	if errPermissionShop != nil {
		internalServerError(w, r, errPermissionShop)
		return
	}
	if !resPermissionShop.IsPermission {
		handleError(w, r, errors.New("not permisson"), 401)
		return
	}

	isPermission, errCheck := c.productService.CheckPermissionOfReformist(profileId, productId)
	if errCheck != nil {
		handleError(w, r, errCheck, 400)
		return
	}
	if !isPermission {
		handleError(w, r, errors.New("not permission"), 401)
		return
	}
	//

	// Update product
	newProduct, err := c.productService.UpdateProduct(product.InfoProduct)
	if err != nil {
		internalServerError(w, r, err)
		return
	}

	// Handle file
	if len(product.ListFileIdDeletes) > 0 {
		_, errDeleteFile := c.clientFileGRPC.DeleteFile(context.Background(), &proto.DeleteFileReq{
			Ids: product.ListFileIdDeletes,
		})
		if errDeleteFile != nil {
			internalServerError(w, r, errDeleteFile)
			return
		}
	}
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

	res := Response{
		Data:    newProduct,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func (c *productController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var product service.DeleyeProductPayload

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		badRequest(w, r, err)
		return
	}

	// Check permission action with product
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")[1]
	mapDataRequest, errMapData := c.utils.JwtDecode(tokenString)
	if errMapData != nil {
		internalServerError(w, r, errMapData)
		return
	}
	profileId := uint(mapDataRequest["profile_id"].(float64))

	resPermissionShop, errPermissionShop := c.clientShopGRPC.CheckShopPermission(
		context.Background(),
		&proto.CheckShopPermissionReq{
			ShopId:    uint64(product.ShopId),
			ProfileId: uint64(profileId),
		},
	)
	if errPermissionShop != nil {
		internalServerError(w, r, errPermissionShop)
		return
	}
	if !resPermissionShop.IsPermission {
		handleError(w, r, errors.New("not permisson"), 401)
		return
	}

	isPermission, errCheck := c.productService.CheckPermissionOfReformist(profileId, product.ProductId)
	if errCheck != nil {
		handleError(w, r, errCheck, 400)
		return
	}
	if !isPermission {
		handleError(w, r, errors.New("not permission"), 401)
		return
	}
	//

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
		productService: service.NewProductService(),
		clientShopGRPC: proto.NewShopServiceClient(config.GetConnShopGRPC()),
		clientFileGRPC: proto.NewFileServiceClient(config.GetConnFileGrpc()),
		utils:          utils.NewJwtUtils(),
	}
}
