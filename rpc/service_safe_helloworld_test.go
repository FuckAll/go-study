package rpc

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"testing"
	"time"
)

const HelloServiceName = "path/to/pkg.HelloService"

type IHelloService interface {
	Hello(request string, reply *string) error
	World(request string, reply *string) error
}

var _ IHelloService = (*HelloServiceImpl)(nil)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}
func (p *HelloServiceImpl) World(request string, reply *string) error {
	*reply = "world:" + request
	return nil
}

// Ignore 虽然是HelloServiceImpl的实现方法，但是由于并没有在接口中，所以不能对外暴露
func (p *HelloServiceImpl) Ignore(request string, reply *string) error {
	*reply = "ignore:" + request
	return nil
}

func RegisterHelloService(service IHelloService) error {
	return rpc.RegisterName(HelloServiceName, service)
}

func TestHelloServiceImpl(t *testing.T) {
	err := RegisterHelloService(&HelloServiceImpl{})
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		rpc.ServeConn(conn)
	}
}

func TestHelloServiceImplClient(t *testing.T) {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	for i := 0; i < 10; i++ {
		var reply string
		err = client.Call(HelloServiceName+".Hello", "hello", &reply)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(reply)

		var world string
		err = client.Call(HelloServiceName+".World", "world", &world)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(world)
		var ignore string
		err = client.Call(HelloServiceName+".Ignore", "ignore", &world)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(ignore)
		<-time.NewTicker(time.Second).C
	}
}
