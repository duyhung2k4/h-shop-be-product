package controller

import (
	"app/service"
	"app/utils"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/go-chi/render"
)

type productController struct {
	productService service.ProductService
	utils          utils.JwtUtils
}

type ProductController interface {
	CreateProduct(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
	DeleteProduct(w http.ResponseWriter, r *http.Request)
}

func (c *productController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		badRequest(w, r, err)
		return
	}

	newProduct, err := c.productService.CreateProduct(product)
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

	isUpdated, errCheck := c.productService.CheckPermissionOfReformist(profileId, productId)
	if errCheck != nil {
		handleError(w, r, errCheck, 400)
		return
	}

	if !isUpdated {
		handleError(w, r, errors.New("Not permission"), 401)
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

	productID := product["_id"].(string)
	err := c.productService.DeleteProduct(productID)
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
		utils:          utils.NewJwtUtils(),
	}
}
