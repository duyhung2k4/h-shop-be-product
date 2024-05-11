package service

import (
	"app/config"
	"app/grpc/proto"
	"app/model"
	"app/utils"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type productService struct {
	db             *mongo.Database
	clientShopGRPC proto.ShopServiceClient
	jwtUtils       utils.JwtUtils
	redisUtils     utils.RedisUtils
}

type ProductService interface {
	GetProductById(productId string) (map[string]interface{}, error)
	CreateProduct(product map[string]interface{}) (map[string]interface{}, error)
	UpdateProduct(product map[string]interface{}) (map[string]interface{}, error)
	DeleteProduct(productId string) error

	CheckPermissionShop(shopId uint, tokenString string) (*bool, error)
	CheckPermissionProduct(productId string, tokenString string) (*bool, error)

	checkPermissionOfReformist(profileId uint, productId string) (bool, error)
}

func (s *productService) GetProductById(productId string) (map[string]interface{}, error) {
	var product map[string]interface{}

	convertToObjectId, errConvert := primitive.ObjectIDFromHex(productId)
	if errConvert != nil {
		return model.MapDataEmpty, errConvert
	}

	filter := bson.M{
		"_id": convertToObjectId,
	}

	err := s.db.
		Collection(string(model.PRODUCT)).
		FindOne(context.Background(), filter).
		Decode(&product)

	if err != nil {
		return model.MapDataEmpty, err
	}

	product["_id"] = product["_id"].(primitive.ObjectID).Hex()
	if err := s.redisUtils.Cache(product["_id"].(string), product); err != nil {
		return model.MapDataEmpty, err
	}

	return product, nil
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

	newProduct["_id"] = newProduct["_id"].(primitive.ObjectID).Hex()
	s.redisUtils.Cache(newProduct["_id"].(string), newProduct)

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

	s.redisUtils.Delete(productId)

	return nil
}

func (s *productService) CheckPermissionShop(shopId uint, tokenString string) (*bool, error) {
	mapDataRequest, errMapData := s.jwtUtils.JwtDecode(tokenString)
	if errMapData != nil {
		return nil, errMapData
	}
	profileId := uint(mapDataRequest["profile_id"].(float64))

	resPermissionShop, errPermissionShop := s.clientShopGRPC.CheckShopPermission(
		context.Background(),
		&proto.CheckShopPermissionReq{
			ShopId:    uint64(shopId),
			ProfileId: uint64(profileId),
		},
	)
	if errPermissionShop != nil {
		return nil, errPermissionShop
	}
	if !resPermissionShop.IsPermission {
		return &model.FALSE_VALUE, nil
	}

	return &model.TRUE_VALUE, nil
}

func (s *productService) CheckPermissionProduct(productId string, tokenString string) (*bool, error) {
	mapDataRequest, errMapData := s.jwtUtils.JwtDecode(tokenString)
	if errMapData != nil {
		return nil, errMapData
	}
	profileId := uint(mapDataRequest["profile_id"].(float64))

	isPermission, errCheck := s.checkPermissionOfReformist(profileId, productId)
	if errCheck != nil {
		return nil, errCheck
	}
	if !isPermission {
		return &model.FALSE_VALUE, nil
	}
	return &model.TRUE_VALUE, nil
}

func (s *productService) checkPermissionOfReformist(profileId uint, productId string) (bool, error) {
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

	if uint(product["profileId"].(int64)) != profileId {
		return false, nil
	}

	return true, nil
}

func NewProductService() ProductService {
	return &productService{
		db:             config.GetDB(),
		clientShopGRPC: proto.NewShopServiceClient(config.GetConnShopGRPC()),
		jwtUtils:       utils.NewJwtUtils(),
		redisUtils:     utils.NewUtilsRedis(),
	}
}
