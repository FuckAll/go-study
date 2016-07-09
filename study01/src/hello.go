package main

import (
	"fmt"
	"unsafe"
)

// "net/http"

const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(b)
)
const (
	d byte = 100
	e int  = 1e10
)

const (
	Sunday             = 1
	Monday             = 2
	Tuesday, Wednesday = iota, iota
)

const (
	A, B = iota, iota << 10
	C, D
)

const (
	_        = iota
	KB int64 = 1 << (10 * iota)
	MB
	GB
	TB
)

type Color int

const (
	Black Color = iota
	Red
	Blue
)

func wonder(c Color) {

}
func main() {
	Ppointer()

}

func Pointer() {
	x := 0x12345678
	p := unsafe.Pointer(&x)
	n := (*[4]byte)(p)
	for i := 0; i < len(n); i++ {
		fmt.Printf("%X", n[i])

	}
}

func Ppointer() {
	d := struct {
		s string
		x int
	}{"abc", 100}
	p := uintptr(unsafe.Pointer(&d))
	p += unsafe.Offsetof(d.x)
	fmt.Println(p)
	p2 := unsafe.Pointer(p)
	fmt.Println(p2)
	px := (*int)(p2)
	*px = 200
	fmt.Printf("%#v\n", d)
}
