package rpc

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"testing"
	"time"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func TestServer(t *testing.T) {
	err := rpc.RegisterName("HelloService", new(HelloService))
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	//for {
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}
	rpc.ServeConn(conn)
	//}
}

func TestClient(t *testing.T) {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	for i := 0; i < 10; i++ {
		var reply string
		err = client.Call("HelloService.Hello", "hello", &reply)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(reply)
		<-time.NewTicker(time.Second).C
	}
}
