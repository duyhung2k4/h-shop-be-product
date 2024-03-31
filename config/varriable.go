package config

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

const (
	APP_PORT             = "APP_PORT"
	APP_NAME             = "APP_NAME"
	DB_HOST              = "DB_HOST"
	DB_PORT              = "DB_PORT"
	DB_WRITE_NAME        = "DB_WRITE_NAME"
	DB_PASSWORD          = "DB_PASSWORD"
	DB_USER              = "DB_USER"
	URL_REDIS            = "URL_REDIS"
	HOST                 = "HOST"
	MONGO_DB_WRITE_LOCAL = "MONGO_DB_WRITE_LOCAL"
)

var (
	appPort string
	// appName string
	// dbHost  string
	// dbPort     string
	dbWriteName string
	// dbPassword        string
	// dbUser            string
	urlRedis          string
	host              string
	mongodbWriteLocal string

	dbWrite *mongo.Database
	rdb     *redis.Client
	jwt     *jwtauth.JWTAuth

	clientProfile *grpc.ClientConn
)
