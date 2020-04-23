//
//  circular_queue.go
//
//  Created by izgnod on 17/9/30.
//  Copyright © 2017年 izgnod. All rights reserved.
//

/**
 *  使用静态数组来实现循环队列
 */

package queue

import (
	"log"
	"sync"
)

// QueueSize queue's max size
const QueueSize = 50

// Queue struct
type Queue struct {
	tail  int // 写
	head  int // 读
	queue [QueueSize]interface{}
	mu    *sync.Mutex
}

// NewQueue new queue
func NewQueue() *Queue {
	q := new(Queue)
	q.tail, q.head = 0, 0
	q.mu = new(sync.Mutex)
	return q
}

// Enqueue 入队列
func (q *Queue) Enqueue(a interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.isFull() {
		// log.Println("队列已满，无法插入数据:", a)
		return
	}
	q.queue[q.tail] = a
	q.tail = (q.tail + 1) % QueueSize
}

// Dequeue 出队列
func (q *Queue) Dequeue() (ret interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.isEmpty() {
		log.Println("队列为空，无法出队列")
		return
	}

	ret = q.queue[q.head]
	q.queue[q.head] = nil
	q.head = (q.head + 1) % QueueSize
	return
}

// isFull
/*
(tail+1)%QueueSize == head
*/

func (q *Queue) isFull() bool {
	// 写的追上读的
	if (q.tail+1)%QueueSize == q.head {
		return true
	}
	return false
}

// isEmpty
/*
tail == head
*/
func (q *Queue) isEmpty() bool {
	if q.tail == q.head {
		return true
	}
	return false
}
