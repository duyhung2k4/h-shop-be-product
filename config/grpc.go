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

	clientProfile, err = grpc.Dial(host+":20002", grpc.WithTransportCredentials(creds))
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
