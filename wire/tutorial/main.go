package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type Message string

func NewMessage(phrase string, ctx context.Context) Message {
	return Message(phrase)
}

type Greeter struct {
	message Message
	grumpy  bool
}

func NewGreeter(m Message, ctx context.Context, good string) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{message: m, grumpy: grumpy}
}

func (g Greeter) Greet() Message {
	if g.grumpy {
		return Message("Go away!")
	}
	return g.message
}

type Event struct {
	greeter Greeter
}

func NewEvent(g Greeter) (Event, error) {
	if g.grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{greeter: g}, nil
}

func (e Event) Start() {
	msg := e.greeter.Greet()
	fmt.Println(msg)
}

func main() {
	//event, err := InitializeEvent("hi there")
	//event, err := InitializeEvent()
	//if err != nil {
	//	fmt.Printf("failed to create event:%s\n", err)
	//	os.Exit(2)
	//}
	//event.Start()
}
