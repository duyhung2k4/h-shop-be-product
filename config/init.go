package config

import (
	"flag"
	"fmt"
	"log"

	"github.com/go-chi/jwtauth/v5"
)

func init() {
	loadEnv()

	var migrate bool = false
	flag.BoolVar(&migrate, "db", true, "Migrate Database?")
	var mongoDBURL string
	flag.StringVar(&mongoDBURL, "host", "localhost", "")
	flag.Parse()

	mongodbLocal = fmt.Sprintf("mongodb://%s:27017", mongoDBURL)
	jwt = jwtauth.New("HS256", []byte("h-shop"), nil)

	if err := connectMongoDB(migrate); err != nil {
		log.Fatalf("Error connect MongoDB : %v", err)
	}
	connectRedis()
	connectGPRCServerShop()
	connectGRPCServerFile()
	connectGRPCServerWarehouse()
	connectElastic()
	connectRabbitMQ()
}
