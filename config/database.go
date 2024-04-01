package config

import (
	"app/model"
	"context"
	"log"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectMongoDB(migrate bool) error {
	var err error
	var client *mongo.Client

	dns := mongodbLocal
	opts := options.Client().ApplyURI(dns)
	client, err = mongo.Connect(context.Background(), opts)

	if err != nil {
		log.Fatalln(err)
	}

	db = client.Database(dbName)

	if migrate {
		// Create Collection
		db.CreateCollection(context.Background(), string(model.PRODUCT))
		db.CreateCollection(context.Background(), string(model.SALE))
	}

	return nil
}

func connectRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     urlRedis,
		Password: "",
		DB:       0,
	})
}
