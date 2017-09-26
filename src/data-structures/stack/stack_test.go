package stack

import "testing"
import "fmt"
import "time"

func TestGood(t *testing.T) {
	stack := NewStack()
	fmt.Printf("%+v\n", stack)

	fmt.Printf("%+v\n", stack.Len())
	stack.Push("string")

	for i := 0; i < 10; i++ {
		go func(input interface{}) {
			stack.Push(input)
		}(i)
	}
	time.Sleep(3 * time.Second)

	fmt.Printf("%+v\n", stack)

	for i := 0; i < 10; i++ {
		fmt.Printf("%+v\n", stack.Pop())
	}

	fmt.Printf("%+v\n", stack.Peek())

	// time.Sleep(3 * time.Second)
	fmt.Printf("%+v\n", stack.Len())

}
