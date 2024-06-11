package controller

import "net/http"

type typeInWarehouseController struct{}

type TypeInWarehouseController interface {
	InsertTypeInWarehouse(w http.ResponseWriter, r *http.Request)
}
