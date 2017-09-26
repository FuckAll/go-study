package queue

import "testing"

func TestQueue(t *testing.T) {
	q := NewQueue()
	// fmt.Printf("%+v\n", q)

	// fmt.Println(q.Len())
	q.Push("string")

	q.Shift()
	q.Shift()
	q.Show()
}
