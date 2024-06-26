package config

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-chi/jwtauth/v5"
	"github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func GetDB() *mongo.Database {
	return db
}

func GetRDB() *redis.Client {
	return rdb
}

func GetAppPort() string {
	return appPort
}

func GetJWT() *jwtauth.JWTAuth {
	return jwt
}

func GetConnShopGRPC() *grpc.ClientConn {
	return clientShop
}

func GetConnFileGrpc() *grpc.ClientConn {
	return clientFile
}

func GetConnWarehouseGrpc() *grpc.ClientConn {
	return clientWarehouse
}

func GetHost() string {
	return host
}

func GetElastic() *elasticsearch.Client {
	return es
}

func GetRabbitChannel() *amqp091.Channel {
	return rabbitChannel
}
