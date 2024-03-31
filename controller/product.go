package controller

import (
	"app/service"
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
)

type productController struct {
	productService service.ProductService
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

func (c *productController) UpdateProduct(w http.ResponseWriter, r *http.Request) {}

func (c *productController) DeleteProduct(w http.ResponseWriter, r *http.Request) {}

func NewProductController() ProductController {
	return &productController{
		productService: service.NewProductService(),
	}
}
