package controller

import "net/http"

type typeInWarehouseController struct{}

type TypeInWarehouseController interface {
	InsertTypeInWarehouse(w http.ResponseWriter, r *http.Request)
}

func (c *typeInWarehouseController) InsertTypeInWarehouse(w http.ResponseWriter, r *http.Request) {}

func NewTypeInWarehouseController() TypeInWarehouseController {
	return &typeInWarehouseController{}
}
