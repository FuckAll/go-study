package main

import (
	"fmt"
)

func main() {
	e1 := (interface{})(nil)
	e2 := (*error)(nil)
	_ = e1 == e2
	fmt.Println(e1 == e2)
}
