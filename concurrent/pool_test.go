package concurrent

import (
	"sync"
	"testing"
	"time"
)

//func Test_Pool(t *testing.T) {
//	pool, _ := NewPool(func() (closer io.Closer, e error) {
//		log.Println("New Closer")
//		return &MyCloseFunc{}, nil
//	}, 10)
//
//	var wg sync.WaitGroup
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		for i := 0; i < 100; i++ {
//			closer, err := pool.Acquire()
//			if err != nil {
//				t.Log(err)
//				return
//			}
//			closeFunc := closer.(*MyCloseFunc)
//			closeFunc.Run()
//			pool.Release(closer)
//			time.Sleep(time.Millisecond * time.Duration(i*10))
//		}
//	}()
//
//	go func() {
//		defer wg.Done()
//		for i := 0; i < 100; i++ {
//			closer, err := pool.Acquire()
//			if err != nil {
//				t.Log(err)
//				return
//			}
//			closeFunc := closer.(*MyCloseFunc)
//			closeFunc.Run()
//			pool.Release(closer)
//		}
//	}()
//
//	wg.Wait()
//
//	pool.Close()
//}
//
//type MyCloseFunc struct {
//}
//
//func (m *MyCloseFunc) Close() error {
//	log.Println("Close")
//	return nil
//}
//
//func (m *MyCloseFunc) Run() error {
//	log.Println("Run")
//	return nil
//}

func Test_WaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Done()
	wg.Wait()

	time.Sleep(time.Second * 10000)
}
