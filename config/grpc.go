package config

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func connectGPRC() {
	var errProfile error

	creds, errKey := credentials.NewClientTLSFromFile("keys/server-shop/public.pem", "localhost")
	if errKey != nil {
		log.Fatalln(errKey)
	}

	clientProfile, errProfile = grpc.Dial(host+":20002", grpc.WithTransportCredentials(creds))
	if errProfile != nil {
		log.Fatalln(errProfile)
	}
}
