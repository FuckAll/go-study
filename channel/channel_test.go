package channel

import (
	"log"
	"sync"
	"testing"
	"time"
)

func Test_CloseChan(t *testing.T) {
	// 无论buffer长度是多少，关闭之后的chan一样可以取出数据
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

	// 关闭之后的chan依旧可以我取出数据
	<-ch

	wg.Wait()
}
