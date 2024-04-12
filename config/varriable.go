package config

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

const (
	APP_PORT       = "APP_PORT"
	APP_NAME       = "APP_NAME"
	DB_HOST        = "DB_HOST"
	DB_PORT        = "DB_PORT"
	DB_NAME        = "DB_NAME"
	DB_PASSWORD    = "DB_PASSWORD"
	DB_USER        = "DB_USER"
	URL_REDIS      = "URL_REDIS"
	HOST           = "HOST"
	MONGO_DB_LOCAL = "MONGO_DB_LOCAL"
)

var (
	appPort string
	// appName string
	// dbHost  string
	// dbPort     string
	dbName string
	// dbPassword        string
	// dbUser            string
	urlRedis     string
	host         string
	mongodbLocal string

	db  *mongo.Database
	rdb *redis.Client
	jwt *jwtauth.JWTAuth

	clientShop *grpc.ClientConn
	clientFile *grpc.ClientConn
)
