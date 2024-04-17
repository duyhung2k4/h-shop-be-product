package config

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func connectGPRCServerShop() {
	var err error

	creds, errKey := credentials.NewClientTLSFromFile("keys/server-shop/public.pem", "localhost")
	if errKey != nil {
		log.Fatalln(errKey)
	}

	clientShop, err = grpc.Dial(host+":20002", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalln(err)
	}
}

func connectGRPCServerFile() {
	var err error

	creds, errKey := credentials.NewClientTLSFromFile("keys/server-file/public.pem", "localhost")
	if errKey != nil {
		log.Fatalln(errKey)
	}

	clientFile, err = grpc.Dial(host+":20004", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalln(err)
	}
}

func connectGRPCServerWarehouse() {
	var err error

	creds, errKey := credentials.NewClientTLSFromFile("keys/server-warehouse/public.pem", "localhost")
	if errKey != nil {
		log.Fatalln(errKey)
	}

	clientWarehouse, err = grpc.Dial(host+":20005", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalln(err)
	}
}
