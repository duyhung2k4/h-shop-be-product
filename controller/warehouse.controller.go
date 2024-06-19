package controller

import (
	"app/config"
	"app/dto/request"
	"app/grpc/proto"
	"app/model"
	"app/utils"
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

type warehouseController struct {
	grpcWarehouse proto.WarehouseServiceClient
	jwtUtils      utils.JwtUtils
	productUtils  utils.ProductUtils
}

type WarehouseController interface {
	GetWarehouseByProductId(w http.ResponseWriter, r *http.Request)
	UpdateWarehouse(w http.ResponseWriter, r *http.Request)
}

func (c *warehouseController) UpdateWarehouse(w http.ResponseWriter, r *http.Request) {
	var payload request.UpdateWarehouseReq
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
	isPermission, errPermission := c.productUtils.CheckPermissionProduct(payload.ProductId, uint64(profileId))
	if errPermission != nil {
		internalServerError(w, r, errPermission)
		return
	}

	if isPermission == &model.FALSE_VALUE {
		internalServerError(w, r, errors.New("not permission"))
		return
	}

	result, err := c.grpcWarehouse.Update(context.Background(), &proto.UpdateWarehouseReq{
		Id:    uint64(payload.Id),
		Count: uint64(payload.Count),
	})

	if err != nil {
		internalServerError(w, r, err)
		return
	}

	res := Response{
		Data:    result,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func (c *warehouseController) GetWarehouseByProductId(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	productId := params.Get("id")
	if productId == "" {
		badRequest(w, r, errors.New("id cannot be empty"))
		return
	}

	result, err := c.grpcWarehouse.GetWarehouseByProductId(context.Background(), &proto.GetWarehouseByProductIdReq{
		ProductId: productId,
	})

	if err != nil {
		internalServerError(w, r, err)
		return
	}

	res := Response{
		Data:    result,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func NewWarehouseController() WarehouseController {
	return &warehouseController{
		grpcWarehouse: proto.NewWarehouseServiceClient(config.GetConnWarehouseGrpc()),
		jwtUtils:      utils.NewJwtUtils(),
		productUtils:  utils.NewProductUtils(),
	}
}
