package config

import (
	"app/model"
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectMongoDBWrite(migrate bool) error {
	var err error
	var client *mongo.Client

	dns := mongodbWriteLocal
	opts := options.Client().ApplyURI(dns)
	client, err = mongo.Connect(context.Background(), opts)

	nodes, _ := client.ListDatabaseNames(context.Background(), bson.M{})
	fmt.Println("Replica set nodes:", nodes)

	if err != nil {
		log.Fatalln(err)
	}

	dbWrite = client.Database(dbWriteName)

	if migrate {
		// Create Collection
		dbWrite.CreateCollection(context.Background(), string(model.PRODUCT))
		dbWrite.CreateCollection(context.Background(), string(model.SALE))
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
