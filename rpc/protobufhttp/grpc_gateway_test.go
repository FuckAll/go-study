package protobufhttp

import (
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"testing"
)

/*
// protoc -I/usr/local/include -I. \
//  -I$GOPATH/src \
//  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
//  --go_out=plugins=grpc:. \
//  hello.proto
//
//protoc -I/usr/local/include -I. \
//  -I$GOPATH/src \
//  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
//  --grpc-gateway_out=logtostderr=true:. \
//  hello.proto
//
//protoc -I/usr/local/include -I. \
//  -I$GOPATH/src \
//  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
//  --swagger_out=logtostderr=true:. \
//  hello.proto
	先要生成grpc stub
    然后生成grpc-gateway
	最后可以生成swagger
*/
func TestGRPCGateWay(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := RegisterYourServiceHandlerFromEndpoint(ctx, mux, "localhost:1234", opts)
	if err != nil {
		t.Error(err)
		return
	}
	panic(http.ListenAndServe(":8080", mux))
}

func TestRunGRPCService(t *testing.T) {
	server := grpc.NewServer()
	RegisterYourServiceServer(server, new(HelloServiceImpl))
	reflection.Register(server)
	listener, err := net.Listen("tcp", "localhost:1234", )
	if err != nil {
		t.Error(err)
		return
	}
	server.Serve(listener)
}
