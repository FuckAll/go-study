package main

import (
	"errors"
	"fmt"
)

func main() {
	test7()

}
func test() {
	defer func() {
		fmt.Println(recover())
	}()

	defer func() {
		panic("defer panic")
	}()
	panic("test panic")
}

func test1() {
	defer func() {
		fmt.Println("goods")
		fmt.Println(recover())
	}()
	defer fmt.Println(recover())
	defer func() {
		func() {
			fmt.Println("defer inner")
			fmt.Println(recover())
		}()
	}()
	panic("test panic")
}

func expect() {
	fmt.Println("wonder")
	recover()

}

func test2() {
	defer expect()
	panic("test panic")
}

func test3(x, y int) {
	var z int
	func() {
		defer func() {
			if recover() != nil {
				z = 0
			}
		}()
		z = x / y
		return
	}()
	fmt.Println("x/y=", z)
}

var ErrDivByZero = errors.New("division by zero")

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("test error")
	}
	return x / y, nil
}

func test4() {
	a := [3]int{1, 2}
	b := [...]int{1, 2, 3, 4}
	fmt.Println(a, b)
	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10},
		{"user2", 20},
	}
	fmt.Println(d)
}

func test5(x []int) {
	fmt.Printf("x: %p\n", x)
	x[1] = 1000

}

func maintest() {
	a := make([]int, 10, 10)

	fmt.Printf("a:%p\n", a)
	test5(a)
	fmt.Println(a)

	data := [...]int{0, 1, 2, 3, 4, 5}
	s := data[2:4]
	s[0] += 100
	s[1] += 200
	fmt.Println(s)
	fmt.Println(data)

	fmt.Println("----------------")
	s1 := []int{0, 1, 2, 3, 8: 100}
	fmt.Println(s1, len(s1), cap(s1))

}

func test6() {
	s := [3]int{1, 2, 3}
	fmt.Println(s)
	fmt.Printf("%p", &s[1])
	p := &s[2]
	*p += 100
	fmt.Printf("\n-----%p\n", s[2])
	fmt.Println(s)
}

func test7() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := data[8:]
	fmt.Println(s)
	s2 := data[:5]
	fmt.Println(s2)

	copy(s2, s)
	fmt.Println(s2)
	fmt.Println(data)

}
