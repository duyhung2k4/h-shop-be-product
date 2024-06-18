package controller

import (
	"app/config"
	"app/dto/request"
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

func NewTypeInWarehouseController() TypeInWarehouseController {
	return &typeInWarehouseController{
		typeInWarehouseGRPC: proto.NewTypeInWarehouseServiceClient(config.GetConnWarehouseGrpc()),
	}
}
