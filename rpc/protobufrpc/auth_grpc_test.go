package protobufrpc

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
	"testing"
)

type Authentication struct {
	UserName string
	Password string
}

func (a *Authentication) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"username": a.UserName, "password": a.Password}, nil
}

func (a *Authentication) RequireTransportSecurity() bool {
	return false

}

type AuthTestServiceImpl struct {
	//auth *AuthTestServiceServer
}

func (a *AuthTestServiceImpl) NeedAuth(ctx context.Context, req *String) (*String, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing credentials")
	}

	var authUserName, authPassword bool
	if userName, ok := md["username"]; ok {
		if userName[0] == "zhangsan" {
			authUserName = true
		}
	}
	if password, ok := md["password"]; ok {
		if password[0] == "123" {
			authPassword = true
		}
	}
	if authUserName && authPassword {
		return &String{Value: "hello:" + req.Value}, nil
	}
	return nil, fmt.Errorf("auth failed")
}

var _ AuthTestServiceServer = (*AuthTestServiceImpl)(nil)

func TestAuthRPCService(t *testing.T) {
	server := grpc.NewServer()
	RegisterAuthTestServiceServer(server, new(AuthTestServiceImpl))
	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		t.Error(err)
	}
	panic(server.Serve(listener))
}

func TestAuthRPCClient(t *testing.T) {
	auth := &Authentication{
		UserName: "zhangsan",
		Password: "123",
	}
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure(), grpc.WithPerRPCCredentials(auth))
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.Close()
	client := NewAuthTestServiceClient(conn)
	needAuth, err := client.NeedAuth(context.Background(), &String{Value: "hello"})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(needAuth)
}
