package algorithms

import "testing"
import "fmt"

func TestChess(t *testing.T) {
	// Chess()
}

func TestCheck(t *testing.T) {
	CheckBigEndian()
}

func TestBigIntDiv(t *testing.T) {
	fmt.Println(BigIntDiv("100000002302030203023", 6938492))
	fmt.Println(BigIntDiv("135", 18))
}

func TestFindNumber(t *testing.T) {
	findNumber([][]int{{1, 2, 3, 4}, {2, 3, 4, 5}, {3, 4, 5, 6}, {4, 5, 6, 7}, {5, 6, 7, 8}}, 7)
}

func TestReplaceBlank(t *testing.T) {
	fmt.Println("替代完成空格之后:", replaceBlack("We are good. "))
}

func TestFibonacci(t *testing.T) {
	fmt.Println("斐波那契数列:", fibs, fibonacci(4))
}

func TestNumberOfOne(t *testing.T) {
	// fmt.Println("cecece:", 3>>1)
	fmt.Println("1的个数:", numberOfOne(5))
	fmt.Println("1的个数:", numberOfOne2(1))
}

func TestReorderOdd(t *testing.T) {
	fmt.Println(reorderOdd([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}
