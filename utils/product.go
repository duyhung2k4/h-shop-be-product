package utils

import (
	"app/config"
	"app/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type productUtils struct {
	db *mongo.Database
}

type ProductUtils interface {
	CheckPermissionProduct(id string, profileId uint64) (*bool, error)
}

func (u *productUtils) CheckPermissionProduct(id string, profileId uint64) (*bool, error) {
	var product map[string]interface{}

	convertToObjectId, errConvert := primitive.ObjectIDFromHex(id)
	if errConvert != nil {
		return nil, errConvert
	}

	filter := bson.M{
		"_id":       convertToObjectId,
		"profileId": profileId,
	}

	err := u.db.
		Collection(string(model.PRODUCT)).
		FindOne(context.Background(), filter).
		Decode(&product)

	if err != nil {
		return nil, err
	}

	if product["_id"] == nil {
		return &model.FALSE_VALUE, nil
	}

	return &model.TRUE_VALUE, nil
}

func NewProductUtils() ProductUtils {
	return &productUtils{
		db: config.GetDB(),
	}
}
