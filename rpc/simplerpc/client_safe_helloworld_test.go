package simplerpc

import (
	"fmt"
	"log"
	"net/rpc"
	"testing"
	"time"
)

var _ IHelloService = (*HelloClientImpl)(nil)

type HelloClientImpl struct {
	*rpc.Client
}

func (p *HelloClientImpl) Hello(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}
func (p *HelloClientImpl) World(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".World", request, reply)
}

// Ignore 虽然是HelloServiceImpl的实现方法，但是由于并没有在接口中，所以不能对外暴露
func (p *HelloClientImpl) Ignore(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Ignore", request, reply)
}
func DialHelloService(network, address string) *HelloClientImpl {
	client, err := rpc.Dial(network, address)
	if err != nil {
		panic(err)
	}
	return &HelloClientImpl{Client: client}
}

func TestSafeHelloClientImpl(t *testing.T) {
	client := DialHelloService("tcp", "localhost:1234")
	for i := 0; i < 10; i++ {
		var reply string
		err := client.Hello("hello", &reply)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(reply)

		var world string
		err = client.World("world", &world)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(world)

		var ignore string
		err = client.Ignore("ignore", &world)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(ignore)
		times := make(chan time.Time, 10)
		times <- time.Time{}
		//times <- time.NewTicker(time.Second / 10).C
		//x :=
		times <- <-time.NewTicker(time.Second / 10).C
		//fmt.Printf(times)
	}
}
