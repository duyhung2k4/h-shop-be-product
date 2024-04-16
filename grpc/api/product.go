package api

import (
	"app/config"
	"app/grpc/proto"
	"app/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type productGrpc struct {
	db *mongo.Database
	proto.UnsafeProductServiceServer
}

func (g *productGrpc) GetProductByListId(ctx context.Context, req *proto.GetProductByListIdReq) (*proto.GetProductByListIdRes, error) {
	var products []ProductQuery
	listObjId := []primitive.ObjectID{}

	for _, item := range req.ProductIds {
		objId, err := primitive.ObjectIDFromHex(item)
		if err != nil {
			return nil, err
		}

		listObjId = append(listObjId, objId)
	}

	cur, err := g.db.Collection(string(model.PRODUCT)).Find(context.Background(), bson.M{
		"_id": bson.M{"$in": listObjId},
	})
	if err != nil {
		return nil, err
	}
	cur.All(context.Background(), &products)

	productsRes := &proto.GetProductByListIdRes{
		Products: []*proto.Product{},
	}

	for _, item := range products {
		productsRes.Products = append(productsRes.Products, &proto.Product{
			Id:    item.ID.Hex(),
			Price: float32(item.Price),
		})
	}

	return productsRes, nil
}

func NewProductGRPC() proto.ProductServiceServer {
	return &productGrpc{
		db: config.GetDB(),
	}
}

type ProductQuery struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Price float64
}
