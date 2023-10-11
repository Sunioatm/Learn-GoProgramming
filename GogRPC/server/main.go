package main

import (
	"fmt"
	"log"
	"net"
	"server/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	s := grpc.NewServer()

	listener, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		log.Fatal(err)
	}

	services.RegisterCalculatorServer(s, services.NewCalculatorServer())
	reflection.Register(s)

	fmt.Println("gRPC server listening on port 5000")

	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

}
