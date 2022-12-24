package echo

import (
	context "context"
	pb "echo/gen/api"
	"fmt"
)

type Service struct {
	pb.UnimplementedEchoServiceServer
}

func (*Service) Echo(ctx context.Context, body *pb.EchoRequest) (*pb.EchoResponse, error) {
	fmt.Println("Got request call with message:", body.Message)

	return &pb.EchoResponse{
		Message: body.Message,
	}, nil
}
