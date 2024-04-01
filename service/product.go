package service

import (
	"app/config"
	"app/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type productService struct {
	db *mongo.Database
}

type ProductService interface {
	CheckPermissionOfReformist(profileId uint, productId string) (bool, error)
	CreateProduct(product map[string]interface{}) (map[string]interface{}, error)
	UpdateProduct(product map[string]interface{}) (map[string]interface{}, error)
	DeleteProduct(productId string) error
}

func (s *productService) CheckPermissionOfReformist(profileId uint, productId string) (bool, error) {
	objID, errObjID := primitive.ObjectIDFromHex(productId)
	if errObjID != nil {
		return false, errObjID
	}

	var product map[string]interface{}

	if err := s.db.
		Collection(string(model.PRODUCT)).
		FindOne(context.Background(), bson.M{"_id": objID}).Decode(&product); err != nil {
		return false, err
	}

	if product["_id"] == nil {
		return false, nil
	}

	if product["profileId"] != profileId {
		return false, nil
	}

	return true, nil
}

func (s *productService) CreateProduct(product map[string]interface{}) (map[string]interface{}, error) {
	var newProduct map[string]interface{}

	result, errInsert := s.db.Collection(string(model.PRODUCT)).InsertOne(context.Background(), product)
	if errInsert != nil {
		return map[string]interface{}{}, errInsert
	}

	if err := s.db.Collection(string(model.PRODUCT)).FindOne(
		context.Background(),
		bson.M{"_id": result.InsertedID}).
		Decode(&newProduct); err != nil {
		return map[string]interface{}{}, err
	}

	return newProduct, nil
}

func (s *productService) UpdateProduct(product map[string]interface{}) (map[string]interface{}, error) {
	var newProduct map[string]interface{}

	idString := product["_id"].(string)
	objID, errObjID := primitive.ObjectIDFromHex(idString)
	if errObjID != nil {
		return map[string]interface{}{}, errObjID
	}
	product["_id"] = objID

	_, errUpdate := s.db.Collection(string(model.PRODUCT)).UpdateOne(
		context.Background(),
		bson.M{
			"_id": objID,
		},
		bson.M{
			"$set": product,
		},
	)
	if errUpdate != nil {
		return map[string]interface{}{}, errUpdate
	}

	if err := s.db.Collection(string(model.PRODUCT)).FindOne(
		context.Background(),
		bson.M{"_id": objID},
	).Decode(&newProduct); err != nil {
		return map[string]interface{}{}, err
	}

	return newProduct, nil
}

func (s *productService) DeleteProduct(productId string) error {
	objID, errObjID := primitive.ObjectIDFromHex(productId)
	if errObjID != nil {
		return errObjID
	}

	s.db.Collection(string(model.PRODUCT)).DeleteOne(
		context.Background(),
		bson.M{"_id": objID},
	)

	return nil
}

func NewProductService() ProductService {
	return &productService{
		db: config.GetDB(),
	}
}
