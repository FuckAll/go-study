package main

import (
	"fmt"

	"github.com/FuckAll/go-study/src/pb"
	"github.com/wothing/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address    = "localhost:50051"
	defaltName = "world-----"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}

	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	fmt.Println("wdon")
	fmt.Println(context.Background())
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: defaltName})
	fmt.Println(r)
	if err != nil {
		log.Fatalf("could not greet :%v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
