package algorithms

import (
	"fmt"
	"testing"
)

// 迭代的方式(这个不好理解，x保存的是当前值，y保存的是先前值)
func fib(n int) int {
	x, y := 1, 1
	for i := 2; i < n; i++ {
		x, y = x+y, x
		fmt.Println(x, y)
	}
	return x
}

// 保存当前和之前的两个状态
func fibNg(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	prev, curr := 1, 1
	for i := 3; i <= n; i++ {
		prev, curr = curr, curr+prev
	}
	return curr
}

func TestFib(t *testing.T) {
	//t.Log(fib(1))
	//t.Log(fib(2))
	//t.Log(fib(3))
	//t.Log(fib(4))
	//t.Log(fib(5))
	t.Log(fib(6))
}

func TestFibNg(t *testing.T) {
	t.Log(fibNg(1))
	t.Log(fibNg(2))
	t.Log(fibNg(3))
	t.Log(fibNg(4))
	t.Log(fibNg(5))
	t.Log(fibNg(6))
}
