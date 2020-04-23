package stack

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

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

	// fmt.Printf()
	// defer_call()

}

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func TestSafeMapUserAgent(t *testing.T) {
	count := 1000
	gw := sync.WaitGroup{}
	gw.Add(count * 3)
	u := UserAges{ages: map[string]int{}}
	add := func(i int) {
		u.Add(fmt.Sprintf("user_%d", i), i)
		gw.Done()
	}
	for i := 0; i < count; i++ {
		go add(i)
		go add(i)
	}
	for i := 0; i < count; i++ {
		go func(i int) {
			defer gw.Done()
			u.Get(fmt.Sprintf("user_%d", i))
		}(i)
	}
	gw.Wait()
	fmt.Println("Done")
}

func TestRectangleNum(t *testing.T) {
	var count int
	m, n := 3, 3
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i != j {
				count += (m - i) * (n - j)
			}
		}
	}

	fmt.Println(count)

}

func TestChangeIp(t *testing.T) {
	var i int
walk:
	for {
		time.Sleep(time.Second)
		fmt.Println(i)
		if i == 10 {
			break walk
		}
		i++
	}

}
