package grpc

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func RunServerGRPC() {
	listenerGRPC, err := net.Listen("tcp", ":20004")

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

	log.Fatalln(grpcServer.Serve(listenerGRPC))
}
