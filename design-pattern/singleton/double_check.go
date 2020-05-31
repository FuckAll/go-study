package singleton

import (
	"sync"
)

var mux sync.Mutex

// 加排它锁，这样会导致所有获取实例的线程都要排队
func GetInstanceLock() *Example {
	mux.Lock()
	defer mux.Unlock()
	if instance == nil {
		instance = new(Example)
	}
	return instance
}

// 双重校验，减少多线程并发获取实例的排队问题
func GetInstanceDoubleCheck() *Example {
	if instance == nil {
		mux.Lock()
		defer mux.Unlock()
		if instance == nil {
			instance = &Example{}
		}
	}
	return instance
}

var once sync.Once

func GetInstanceByOnce() *Example {
	once.Do(func() {
		instance = new(Example)
	})
	return instance
}
