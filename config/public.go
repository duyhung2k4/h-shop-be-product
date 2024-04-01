package config

import (
	"github.com/go-chi/jwtauth/v5"
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

func GetConnProfileGRPC() *grpc.ClientConn {
	return clientProfile
}

func GetHost() string {
	return host
}
