package grpc

import (
	"app/grpc/api"
	"app/grpc/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func RunServerGRPC() {
	listenerGRPC, err := net.Listen("tcp", ":20003")

	if err != nil {
		log.Fatalln(listenerGRPC)
	}

	creds, errKey := credentials.NewServerTLSFromFile(
		"keys/server-product/public.pem",
		"keys/server-product/private.pem",
	)

	if errKey != nil {
		log.Fatalln(errKey)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))
	proto.RegisterProductServiceServer(grpcServer, api.NewProductGRPC())

	log.Fatalln(grpcServer.Serve(listenerGRPC))
}
