package concurrent

import (
	"log"
	"testing"
	"time"
)

func TestWorkerPool(t *testing.T) {
	workerPool := NewWorkerPool(10)
	for i := 0; i < 100000; i++ {
		workerPool.Add(new(TestWorker))
	}
	workerPool.ShutDown()
}

type TestWorker struct {
}

func (t *TestWorker) Task() {
	log.Print("Task run")
	<-time.After(time.Second * 10)
}
