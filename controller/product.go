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
	var product CreateProductPayload

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		badRequest(w, r, err)
		return
	}

	log.Println("Product: ", product.InfoProduct)

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

	newProduct, errProduct := c.productService.CreateProduct(product.InfoProduct)
	if errProduct != nil {
		internalServerError(w, r, errProduct)
		return
	}

	dataFiles, errDataFiles := c.clientFileGRPC.InsertFile(context.Background(), &proto.InsertFileReq{
		Data:      product.Files,
		TypeModel: string(model.PRODUCT),
		ProductId: newProduct["_id"].(primitive.ObjectID).String(),
	})
	if errDataFiles != nil {
		internalServerError(w, r, errDataFiles)
		return
	}

	newProduct["fileIds"] = dataFiles.FileIds

	res := Response{
		Data:    newProduct,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func (c *productController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		badRequest(w, r, err)
		return
	}

	tokenString := strings.Split(r.Header.Get("Authorization"), " ")[1]
	mapDataRequest, errMapData := c.utils.JwtDecode(tokenString)
	if errMapData != nil {
		internalServerError(w, r, errMapData)
		return
	}
	profileId := uint(mapDataRequest["profile_id"].(float64))
	productId := product["_id"].(string)
	shopId := uint(product["shopId"].(float64))

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

	newProduct, err := c.productService.UpdateProduct(product)
	if err != nil {
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
	var product map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		badRequest(w, r, err)
		return
	}

	tokenString := strings.Split(r.Header.Get("Authorization"), " ")[1]
	mapDataRequest, errMapData := c.utils.JwtDecode(tokenString)
	if errMapData != nil {
		internalServerError(w, r, errMapData)
		return
	}
	profileId := uint(mapDataRequest["profile_id"].(float64))
	productId := product["_id"].(string)
	shopId := uint(product["shopId"].(float64))

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

	err := c.productService.DeleteProduct(productId)
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
		clientShopGRPC: proto.NewShopServiceClient(config.GetConnProfileGRPC()),
		clientFileGRPC: proto.NewFileServiceClient(config.GetConnFileGrpc()),
		utils:          utils.NewJwtUtils(),
	}
}

type CreateProductPayload struct {
	InfoProduct map[string]interface{} `json:"infoProduct"`
	Files       [][]byte               `json:"files"`
}
