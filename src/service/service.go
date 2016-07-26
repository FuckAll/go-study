package main

import (
	"net"

	"github.com/FuckAll/go-study/src/pb"

	"github.com/wothing/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct {
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: ("hello" + in.Name)}, nil
}

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}
