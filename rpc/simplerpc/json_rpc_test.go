package simplerpc

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"testing"
	"time"
)

func TestJsonRPCService(t *testing.T) {
	err := RegisterHelloService(&HelloServiceImpl{})
	if err != nil {
		panic(err)
	}

	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		t.Error(err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			t.Error(err)
			break
		}
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

func TestJsonRPCClient(t *testing.T) {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	for i := 0; i < 100; i++ {
		var resp string
		err = client.Call(HelloServiceName+".Hello", "hello", &resp)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(resp)
		<-time.NewTicker(time.Second).C
	}

}
