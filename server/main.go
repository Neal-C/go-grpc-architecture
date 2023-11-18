package main

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (self *Server) Divide(ctx context.Context, in *protocodegen.CalculationRequest) (*protocodegen.CalculationResponse, error) {
	if in.B == 0 {
		return nil, status.Error(codes.InvalidArgument, "cannot divide by zero")
	}
	return &protocodegen.CalculationResponse{
		Result: in.A / in.B,
	}, nil
}

func (self *Server) Sum(ctx context.Context, in *protocodegen.NumbersRequest) (*protocodegen.CalculationResponse, error) {

	var result int64

	for _, n := range in.Numbers {
		result += n
	}

	return &protocodegen.CalculationResponse{
		Result: result,
	}, nil
}

func (self *Server) Sub(ctx context.Context, in *protocodegen.CalculationRequest) (*protocodegen.CalculationResponse, error) {

	return &protocodegen.CalculationResponse{
		Result: in.A - in.B,
	}, nil
}

func (self *Server) Multiply(ctx context.Context, in *protocodegen.CalculationRequest) (*protocodegen.CalculationResponse, error) {
	return &protocodegen.CalculationResponse{
		Result: in.A * in.B,
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
