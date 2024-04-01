package config

import (
	"app/model"
	"context"
	"log"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
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

		// Index
		// Index - Product
		if _, err := db.Collection(string(model.PRODUCT)).
			Indexes().
			CreateMany(context.Background(), []mongo.IndexModel{
				{Keys: bson.M{"name": 1}, Options: options.Index().SetName("idx_name")},
			}); err != nil {
			log.Fatalln(err)
		}
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
