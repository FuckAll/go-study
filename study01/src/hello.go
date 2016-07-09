package main

import (
	"fmt"
	"unsafe"
	// "github.com/urfave/negroni"
	// "net/http"
)

const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(b)
)
const (
	d byte = 100
	e int  = 1e10
)

const {
	Sunday = iota
	Monday 
	Tuesday
}
func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

}
