package context_test

import (
	"context"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func Test_Context(t *testing.T) {

	c1, cal1 := context.WithCancel(context.Background())
	defer cal1()

	c2, cal2 := context.WithCancel(c1)
	defer cal2()

	_ = c2

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, 1500*time.Millisecond)

	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())

	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

func Test_Select(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func(ctx context.Context) {
		defer wg.Done()
		select {
		case a := <-ctx.Done():
			log.Printf("this is a :%d", a)
		}

	}(ctx)

	go func(ctx context.Context) {
		defer wg.Done()
		select {
		case b := <-ctx.Done():
			log.Printf("this is b :%d", b)
		}
	}(ctx)

	time.Sleep(10 * time.Second)
	cancel()
	wg.Wait()
}

func Test_CloseChan(t *testing.T) {
	ch := make(chan struct{}, 10)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func(ch chan struct{}) {
		defer wg.Done()
		select {
		case a := <-ch:
			log.Printf("this is a :%d", a)
		}

	}(ch)

	go func(ch chan struct{}) {
		defer wg.Done()
		select {
		case b := <-ch:
			log.Printf("this is b :%d", b)
		}
	}(ch)

	time.Sleep(time.Second * 3)

	close(ch)
	<-ch
	wg.Wait()
}
