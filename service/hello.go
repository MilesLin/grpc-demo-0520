package service

import (
	"context"
	"grpcdemo/pb"
)

type HelloService struct{}

func (*HelloService) SayHello(c context.Context, h *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hi " + h.Name + ", this message from gRPC server"}, nil
}
