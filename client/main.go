package main

import (
	"context"
	"flag"
	"github.com/Neal-C/go-grpc-architecture/protocodegen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	serverAddr := flag.String("server", "localhost:8080", "the server address in format -> host:port")

	flag.Parse()

	// tlsCredentials := credentials.NewTLS(&tls.Config{InsecureSkipVerify: false})

	grpcDialOptions := []grpc.DialOption{
		// grpc.WithTransportCredentials(tlsCredentials),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	grpcConnection, err := grpc.DialContext(ctx, *serverAddr, grpcDialOptions...)

	if err != nil {
		log.Fatalln("failed to dial grpcServer %s", *serverAddr, err)
	}

	defer grpcConnection.Close()

	grpcClient := protocodegen.NewCalculatorClient(grpcConnection)

	response, err := grpcClient.Sum(ctx, &protocodegen.NumbersRequest{
		Numbers: []int64{10, 10, 10, 42},
	})

	if err != nil {
		log.Fatalln("error sending the &protocodegen.NumbersRequest : ", err)
	}

	log.Println(response)
}
