package concurrent

import "sync"

type Worker interface {
	Task()
}

type WorkerPool struct {
	worker chan Worker
	wg     sync.WaitGroup
}

func NewWorkerPool(maxGoroutines int) *WorkerPool {
	pool := &WorkerPool{
		worker: make(chan Worker),
	}

	pool.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for worker := range pool.worker {
				worker.Task()
			}
			pool.wg.Done()
		}()
	}
	return pool
}

func (p *WorkerPool) Add(worker Worker) {
	p.worker <- worker
}

func (p *WorkerPool) ShutDown() {
	close(p.worker)
	p.wg.Wait()
}
