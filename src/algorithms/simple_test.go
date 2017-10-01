package simple

import "testing"
import "fmt"

func TestChess(t *testing.T) {
	Chess()
}

func TestCheck(t *testing.T) {
	CheckBigEndian()
}

func TestBigIntDiv(t *testing.T) {
	fmt.Println(BigIntDiv("100000002302030203023", 6938492))
	fmt.Println(BigIntDiv("135", 18))
}
