package main

import (
	"flag"
	"fmt"
)

func test(a []int) {
	a[1] = 4
}
func main() {
	var arrlazy = []int{5, 6, 6, 7}
	fmt.Println(arrlazy)
	test(arrlazy)
	fmt.Println(arrlazy)

}
