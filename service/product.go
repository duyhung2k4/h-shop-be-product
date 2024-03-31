package service

import (
	"app/config"
	"app/model"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type productService struct {
	dbWrite *mongo.Database
}

type ProductService interface {
	CreateProduct(product map[string]interface{}) (map[string]interface{}, error)
}

func (s *productService) CreateProduct(product map[string]interface{}) (map[string]interface{}, error) {
	var newProduct map[string]interface{}

	result, errInsert := s.dbWrite.Collection(string(model.PRODUCT)).InsertOne(context.Background(), product)
	if errInsert != nil {
		return map[string]interface{}{}, errInsert
	}

	if err := s.dbWrite.Collection(string(model.PRODUCT)).FindOne(context.Background(), map[string]interface{}{
		"_id": result.InsertedID,
	}).Decode(&newProduct); err != nil {
		return map[string]interface{}{}, err
	}

	return newProduct, nil
}

func NewProductService() ProductService {
	return &productService{
		dbWrite: config.GetDBWrite(),
	}
}
