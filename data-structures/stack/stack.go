package stack

import (
	"sync"
)

// Stack struct
type Stack struct {
	stack []interface{}
	len   int
	mu    *sync.Mutex
}

// NewStack return stack
func NewStack() *Stack {
	stack := new(Stack)
	stack.mu = new(sync.Mutex)
	stack.len = 0
	return stack
}

// Len return stack length
func (s *Stack) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.len
}

// Push push to Stack
func (s *Stack) Push(i interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()

	prepend := make([]interface{}, 1)
	prepend[0] = i
	prepend = append(prepend, s.stack...)
	s.stack = prepend
	s.len++
	return
}

// Pop pop from stack
func (s *Stack) Pop() (el interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()

	el, s.stack = s.stack[0], s.stack[1:]
	s.len--
	return
}

// IsEmpty check stack empty
func (s *Stack) IsEmpty() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.len == 0
}

// Peek retrun first data
func (s *Stack) Peek() interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.stack[0]
}
