package main

import (
	pb "echo/gen/api"
	echo "echo/services"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	fmt.Println("Starting gRPC server...")

	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		panic(err)
	}

	defer lis.Close()

	s := grpc.NewServer()
	echoService := echo.Service{}

	pb.RegisterEchoServiceServer(s, &echoService)

	forever := make(chan bool)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("faild to serve: %v", err)
		}
	}()

	fmt.Println("gRPC server started at port 9000!")

	defer s.Stop()

	<-forever
}
