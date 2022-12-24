package main

import (
	"context"
	pb "echo/gen/api"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial("localhost:9000", opts...)
	if err != nil {
		log.Fatal("Cannot dial grpc server")
	}

	defer conn.Close()

	s := pb.NewEchoServiceClient(conn)

	fmt.Println("Trying to call Echo procedure...")
	body := pb.EchoRequest{
		Message: "ping",
	}

	response, err := s.Echo(context.Background(), &body)

	if err != nil {
		log.Fatalf("Echo grpc call fail: %v", err)
	}

	fmt.Println("Got Echo message:", response.Message)
}
