package queue

import "testing"

func TestQueue(t *testing.T) {
	q := NewQueue()
	q.Push("string")

	q.Shift()
	q.Shift()
	q.Show()
}
