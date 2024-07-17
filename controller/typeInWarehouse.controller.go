package controller

import (
	"app/config"
	"app/dto/request"
	"app/dto/response"
	"app/grpc/proto"
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

type typeInWarehouseController struct {
	typeInWarehouseGRPC proto.TypeInWarehouseServiceClient
}

type TypeInWarehouseController interface {
	GetTypeInWarehouseByProductId(w http.ResponseWriter, r *http.Request)
	InsertTypeInWarehouse(w http.ResponseWriter, r *http.Request)
	UpdateTypeInWarehouse(w http.ResponseWriter, r *http.Request)
	DeleteTypeInWarehouse(w http.ResponseWriter, r *http.Request)
}

func (c *typeInWarehouseController) InsertTypeInWarehouse(w http.ResponseWriter, r *http.Request) {
	var payload request.InsertTypeInWarehouseReq
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		badRequest(w, r, err)
		return
	}

	result, err := c.typeInWarehouseGRPC.Insert(context.Background(), &proto.InsertTypeInWarehouseReq{
		ProductId: payload.ProductId,
		Name:      payload.Name,
		HasTag:    payload.Hastag,
		Price:     float32(*payload.Price),
		Count:     uint64(payload.Count),
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

func (c *typeInWarehouseController) GetTypeInWarehouseByProductId(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	productId := params.Get("id")
	if productId == "" {
		badRequest(w, r, errors.New("id not empty"))
		return
	}

	result, err := c.typeInWarehouseGRPC.GetTypeInWarehouseByProductId(context.Background(), &proto.GetTypeInWarehouseByProductIdReq{
		ProductId: productId,
	})

	if err != nil {
		internalServerError(w, r, err)
		return
	}

	res := Response{
		Data:    result.Data,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func (c *typeInWarehouseController) UpdateTypeInWarehouse(w http.ResponseWriter, r *http.Request) {
	var payload request.UpdateTypeInWarehouseReq
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		badRequest(w, r, err)
		return
	}

	result, errUpdate := c.typeInWarehouseGRPC.Update(context.Background(), &proto.UpdateTypeInWarehouseReq{
		Id:     payload.Id,
		Count:  uint64(payload.Count),
		Price:  float32(*payload.Price),
		Hastag: payload.Hastag,
		Name:   payload.Name,
	})

	if errUpdate != nil {
		internalServerError(w, r, errUpdate)
		return
	}

	data := response.UpdateTypeInWarehouseRes{
		Id:     result.Id,
		Count:  uint(result.Count),
		Name:   result.Name,
		Hastag: result.Hastag,
	}
	if result.Price != 0 {
		price := float64(result.Price)
		data.Price = &price
	}

	res := Response{
		Data:    data,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}
func (c *typeInWarehouseController) DeleteTypeInWarehouse(w http.ResponseWriter, r *http.Request) {
	var payload request.DeleteTypeInWarehouseReq

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		badRequest(w, r, err)
		return
	}

	result, errDelete := c.typeInWarehouseGRPC.Delete(context.Background(), &proto.DeleteTypeInWarehouseReq{
		Id: payload.Id,
	})

	if errDelete != nil {
		internalServerError(w, r, errDelete)
		return
	}

	data := response.DeleteTypeInWarehouseRes{
		Success: result.Success,
	}

	res := Response{
		Data:    data,
		Status:  200,
		Message: "OK",
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func NewTypeInWarehouseController() TypeInWarehouseController {
	return &typeInWarehouseController{
		typeInWarehouseGRPC: proto.NewTypeInWarehouseServiceClient(config.GetConnWarehouseGrpc()),
	}
}
