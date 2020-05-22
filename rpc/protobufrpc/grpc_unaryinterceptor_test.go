package protobufrpc

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"testing"
)

//type UnaryServerInterceptor func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler UnaryHandler) (resp interface{}, err error)
func filter(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Printf("filter:%+v\n", info)
	// info#Service中存放的是interface类型
	if srv, ok := info.Server.(HelloServiceServer); ok {
		log.Printf("srv:%+v, ok:%+v\n", srv, ok)
		if impl, o := srv.(*HelloServiceImpl); o {
			hello, err := impl.Hello(ctx, req.(*String))
			log.Printf("xxxxxxx %+v,  %+v\n", hello, err)
		}
	}

	return handler(ctx, req)
}

func TestGRPCUnaryInterceptor(t *testing.T) {
	server := grpc.NewServer(grpc.UnaryInterceptor(filter))
	RegisterHelloServiceServer(server, new(HelloServiceImpl))
	listen, err := net.Listen("tcp", ":12345")
	if err != nil {
		panic(err)
	}
	panic(server.Serve(listen))
}
