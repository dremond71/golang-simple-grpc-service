package main

import (
	"log"
	"net"

	"fmt"

	"github.com/dremond71/golang-simple-grpc-service/upper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	portNumber := 9000

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port %v: %v", portNumber, err)
	}

	s := upper.Server{}

	grpcServer := grpc.NewServer()

	upper.RegisterUpperServiceServer(grpcServer, &s)

	servingMessage := fmt.Sprintf("Serving on port %v", portNumber)

	// enable reflection, otherwise clients (postman, grpcurl, etc) won't be able
	// to discover what your service provides
	reflection.Register(grpcServer)

	fmt.Println(servingMessage)
	log.Println("INFO: " + servingMessage)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to server gRPC server over port %v: %v", portNumber, err)
	}

}
