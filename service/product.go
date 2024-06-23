package service

import (
	"app/config"
	"app/grpc/proto"
	"app/model"
	"app/utils"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type productService struct {
	db             *mongo.Database
	clientShopGRPC proto.ShopServiceClient
	jwtUtils       utils.JwtUtils
	redisUtils     utils.RedisUtils
}

type ProductService interface {
	GetProductbyProfileId(profileId uint64) ([]map[string]interface{}, error)
	GetProductById(productId string) (map[string]interface{}, error)
	CreateProduct(product map[string]interface{}) (map[string]interface{}, error)
	UpdateProduct(product map[string]interface{}) (map[string]interface{}, error)
	DeleteProduct(productId string) error

	Heart(productId string, profileId uint) error
	IsHeart(productId string, profileId uint) bool
	GetHeart(profileId uint) ([]map[string]interface{}, error)

	Cart(productId string, profileId uint) error
	IsCart(productId string, profileId uint) bool
	GetCart(profileId uint) ([]map[string]interface{}, error)

	CheckPermissionShop(shopId uint, tokenString string) (*bool, error)
	CheckPermissionProduct(productId string, tokenString string) (*bool, error)

	checkPermissionOfReformist(profileId uint, productId string) (bool, error)
}

func (s *productService) GetProductbyProfileId(profileId uint64) ([]map[string]interface{}, error) {
	var products []map[string]interface{}

	filter := bson.M{
		"profileId": profileId,
		"deleteAt":  nil,
	}

	cur, err := s.db.Collection(string(model.PRODUCT)).Find(context.Background(), filter)
	if err != nil {
		return model.ArrMapDataEmpty, err
	}

	if err := cur.All(context.Background(), &products); err != nil {
		return model.ArrMapDataEmpty, err
	}

	return products, nil
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
	profileId := product["profileId"].(uint)
	resultProfile, errShopId := s.clientShopGRPC.GetShopByProfileId(context.Background(), &proto.GetShopByProfileIdReq{ProfileId: uint64(profileId)})

	if errShopId != nil {
		return nil, errShopId
	}
	product["shopId"] = resultProfile.ShopId
	product["createAt"] = time.Now()
	product["updateAt"] = nil
	product["deleteAt"] = nil

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
	var currentProduct map[string]interface{}

	idString := product["_id"].(string)
	objID, errObjID := primitive.ObjectIDFromHex(idString)
	if errObjID != nil {
		return map[string]interface{}{}, errObjID
	}

	if err := s.db.Collection(string(model.PRODUCT)).FindOne(context.Background(), bson.M{"_id": objID}).Decode(&currentProduct); err != nil {
		return nil, err
	}

	product["_id"] = objID
	product["updateAt"] = time.Now()
	product["createAt"] = currentProduct["createAt"]
	product["profileId"] = currentProduct["profileId"]
	product["deleteAt"] = nil

	_, errUpdate := s.db.Collection(string(model.PRODUCT)).ReplaceOne(
		context.Background(),
		bson.M{
			"_id": objID,
		},
		product,
		options.Replace().SetUpsert(true),
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

	s.db.Collection(string(model.PRODUCT)).UpdateOne(
		context.Background(),
		bson.M{
			"_id": objID,
		},
		bson.M{
			"$set": map[string]interface{}{
				"deleteAt": time.Now(),
			},
		},
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

func (s *productService) Heart(productId string, profileId uint) error {
	isHeart := s.IsHeart(productId, profileId)

	heart := map[string]interface{}{
		"productId": productId,
		"profileId": profileId,
	}

	if !isHeart {
		_, errInsert := s.db.Collection(string(model.HEART)).InsertOne(context.Background(), heart)
		return errInsert
	}

	_, err := s.db.Collection(string(model.HEART)).DeleteOne(context.Background(), bson.M{
		"productId": productId,
		"profileId": profileId,
	})

	return err
}

func (s *productService) GetHeart(profileId uint) ([]map[string]interface{}, error) {
	var hearts []map[string]interface{}
	ids := []primitive.ObjectID{}

	filterHeart := bson.M{
		"profileId": profileId,
	}

	cur, err := s.db.Collection(string(model.HEART)).Find(context.Background(), filterHeart)
	if err != nil {
		return model.ArrMapDataEmpty, err
	}

	if err := cur.All(context.Background(), &hearts); err != nil {
		return model.ArrMapDataEmpty, err
	}

	for _, item := range hearts {
		objectID, _ := primitive.ObjectIDFromHex(item["productId"].(string))
		ids = append(ids, objectID)
	}

	var products []map[string]interface{}
	filterProduct := bson.M{"_id": bson.M{"$in": ids}}

	curProduct, errProduct := s.db.Collection(string(model.PRODUCT)).Find(context.Background(), filterProduct)
	if errProduct != nil {
		return model.ArrMapDataEmpty, errProduct
	}

	if err := curProduct.All(context.Background(), &products); err != nil {
		return model.ArrMapDataEmpty, err
	}

	return products, nil
}

func (s *productService) IsHeart(productId string, profileId uint) bool {
	var heart map[string]interface{}
	if err := s.db.
		Collection(string(model.HEART)).
		FindOne(context.Background(), bson.M{
			"profileId": profileId,
			"productId": productId,
		}).Decode(&heart); err != nil {
		return false
	}

	if heart["_id"] == nil {
		return false
	}

	return true
}

func (s *productService) Cart(productId string, profileId uint) error {
	isCart := s.IsCart(productId, profileId)

	cart := map[string]interface{}{
		"productId": productId,
		"profileId": profileId,
	}

	if !isCart {
		_, errInsert := s.db.Collection(string(model.CART)).InsertOne(context.Background(), cart)
		return errInsert
	}

	_, err := s.db.Collection(string(model.CART)).DeleteOne(context.Background(), bson.M{
		"productId": productId,
		"profileId": profileId,
	})

	return err
}

func (s *productService) IsCart(productId string, profileId uint) bool {
	var cart map[string]interface{}
	if err := s.db.
		Collection(string(model.CART)).
		FindOne(context.Background(), bson.M{
			"profileId": profileId,
			"productId": productId,
		}).Decode(&cart); err != nil {
		return false
	}

	if cart["_id"] == nil {
		return false
	}

	return true
}

func (s *productService) GetCart(profileId uint) ([]map[string]interface{}, error) {
	var carts []map[string]interface{}
	ids := []primitive.ObjectID{}

	filterCart := bson.M{
		"profileId": profileId,
	}

	cur, err := s.db.Collection(string(model.CART)).Find(context.Background(), filterCart)
	if err != nil {
		return model.ArrMapDataEmpty, err
	}

	if err := cur.All(context.Background(), &carts); err != nil {
		return model.ArrMapDataEmpty, err
	}
	for _, item := range carts {
		objectID, _ := primitive.ObjectIDFromHex(item["productId"].(string))
		ids = append(ids, objectID)
	}

	var products []map[string]interface{}
	filterProduct := bson.M{"_id": bson.M{"$in": ids}}

	curProduct, errProduct := s.db.Collection(string(model.PRODUCT)).Find(context.Background(), filterProduct)
	if errProduct != nil {
		return model.ArrMapDataEmpty, err
	}

	if err := curProduct.All(context.Background(), &products); err != nil {
		return model.ArrMapDataEmpty, err
	}

	return products, nil
}

func (s *productService) checkPermissionOfReformist(profileId uint, productId string) (bool, error) {
	objID, errObjID := primitive.ObjectIDFromHex(productId)
	if errObjID != nil {
		return false, errObjID
	}

	var product map[string]interface{}

	if err := s.db.
		Collection(string(model.PRODUCT)).
		FindOne(context.Background(), bson.M{
			"_id":       objID,
			"profileId": profileId,
		}).Decode(&product); err != nil {
		return false, err
	}

	if product["_id"] == nil {
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
