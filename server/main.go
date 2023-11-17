package main

import (
	"context"
	"log"
	"net"

	"github.com/Neal-C/go-grpc-architecture/protocodegen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	protocodegen.UnimplementedCalculatorServer
}

func (self *Server) Add(ctx context.Context, in *protocodegen.CalculationRequest) (*protocodegen.CalculationResponse, error) {
	return &protocodegen.CalculationResponse{
		Result: in.A + in.B,
	}, nil
}

func main() {
	tcpListener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalln("failed to initialize the tcp listener :", err)
	}
	defer tcpListener.Close()

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	protocodegen.RegisterCalculatorServer(grpcServer, &Server{})

	if err := grpcServer.Serve(tcpListener); err != nil {
		log.Fatalln("Failed to serve server")
	}

}
