package protobufrpc

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"net"
	"testing"
	"time"
)

type HelloServiceImpl struct {
}

func (h *HelloServiceImpl) Hello(ctx context.Context, req *String) (*String, error) {
	return &String{Value: "hello:" + req.Value}, nil
}

func (h *HelloServiceImpl) Channel(ch HelloService_ChannelServer) error {
	for {
		recv, err := ch.Recv()
		if err != nil {
			if err == io.EOF {
				return nil

			}
			return err
		}
		fmt.Printf("接收到数据:%+v\n", recv)
		reply := &String{Value: "hello:" + recv.GetValue()}
		err = ch.Send(reply)
		if err != nil {
			return err
		}
	}
}

func TestHelloServiceImpl(t *testing.T) {
	server := grpc.NewServer()
	RegisterHelloServiceServer(server, &HelloServiceImpl{})
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		t.Error(err)
	}
	err = server.Serve(listener)
	if err != nil {
		t.Error(err)
	}
}

func TestHelloChannelClient(t *testing.T) {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		t.Error(err)
	}

	defer conn.Close()
	client := NewHelloServiceClient(conn)
	channelClient, err := client.Channel(context.Background())
	if err != nil {
		t.Error(err)
	}

	go func() {
		for {
			err = channelClient.Send(&String{Value: "hello"})
			if err != nil {
				t.Error(err)
				fmt.Println("进入")
			}
			time.Sleep(time.Second)
		}
	}()

	for {
		recv, err := channelClient.Recv()
		if err != nil {
			t.Error(err)
		}
		fmt.Printf("recv:%+v\n", recv)
		time.Sleep(time.Second)
	}
}

func TestHelloClientImpl(t *testing.T) {
	//conn, err := grpc.Dial("localhost:12345", grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:3999", grpc.WithInsecure())
	if err != nil {
		t.Error(err)
	}

	defer conn.Close()
	client := NewHelloServiceClient(conn)
	for {
		hello, err := client.Hello(context.Background(), &String{Value: "张三"})
		if err != nil {
			t.Error(err)
		}
		fmt.Printf("grpc: %+v\n", hello)
		time.Sleep(time.Second)
	}

}
