package queue

import "testing"
import "fmt"

func TestQueueEnqueueAndDequeue(t *testing.T) {
	q := NewQueue()
	// for i := 0; i < 200; i++ {
	// 	go func(a int) {
	// 		q.Enqueue(a)
	// 		// q.Dequeue()
	// 	}(i)
	// }

	for i := 0; i < 200; i++ {
		q.Enqueue(i)
	}

	fmt.Printf("%+v\n", q)

	q.Dequeue()

	fmt.Printf("%+v\n", q)

	q.Enqueue("xxxx")
	fmt.Printf("%+v\n", q)

	q.Dequeue()
	fmt.Printf("%+v\n", q)

	q.Enqueue("yyyxxxx")
	fmt.Printf("%+v\n", q)

	qEmpty := NewQueue()
	qEmpty.Dequeue()

	// q.Enqueue("xxx")
	// fmt.Printf("%+v\n", q)
	// q.Enqueue("yyy")
	// fmt.Printf("%+v\n", q)

	// fmt.Println(q.Dequeue())
	// for i := 0; i < 10; i++ {
	// 	q.Dequeue()
	// 	fmt.Printf("%+v\n", q)
	// }

}
