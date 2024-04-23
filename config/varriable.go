package config

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-chi/jwtauth/v5"
	"github.com/rabbitmq/amqp091-go"
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

	ELASTIC_USER     = "ELASTIC_USER"
	ELASTIC_PASSWORD = "ELASTIC_PASSWORD"
	URL_RABBIT_MQ    = "URL_RABBIT_MQ"
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
	urlRabbitMq  string

	db  *mongo.Database
	rdb *redis.Client
	jwt *jwtauth.JWTAuth

	clientShop      *grpc.ClientConn
	clientFile      *grpc.ClientConn
	clientWarehouse *grpc.ClientConn

	elasticUser     string
	elasticPassword string

	es            *elasticsearch.Client
	rabbitChannel *amqp091.Channel
)
