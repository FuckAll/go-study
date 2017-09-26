package queue

import "sync"
import "fmt"

// Queue struct
type Queue struct {
	queue []interface{}
	len   int
	mu    *sync.Mutex // 队列的锁
}

// NewQueue create a queue
func NewQueue() *Queue {
	q := new(Queue)
	// q.queue = make
	q.mu = new(sync.Mutex)
	return q
}

// Len return length queue
func (q *Queue) Len() int {
	return q.len
}

// Shift delete
func (q *Queue) Shift() interface{} {
	var re interface{}
	if q.len > 0 {
		re, q.queue = q.queue[0], q.queue[1:]
		q.len--
		return re
	}
	return nil
}

// Push insert
func (q *Queue) Push(i interface{}) {
	q.queue = append(q.queue, i)
	q.len++
}

func (q *Queue) isEmpty() bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.len == 0
}

// Show show all data
func (q *Queue) Show() {
	for _, d := range q.queue {
		fmt.Printf("%+v\n", d)
	}
}
