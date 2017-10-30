package tree

import (
	"container/list"
	"fmt"
	"testing"
)

func TestFindPath(t *testing.T) {
	//            10
	//         /      \
	//        5        12
	//       /\
	//      4  7

	b1 := new(Btree)
	b1.value = 10

	b2 := new(Btree)
	b2.value = 5
	b1.left = b2

	b3 := new(Btree)
	b3.value = 12
	b1.right = b3

	b4 := new(Btree)
	b4.value = 4
	b2.left = b4

	b5 := new(Btree)
	b5.value = 7
	b2.right = b5

	// showTree(b1)

	FindPath(b1, 22)
}

func showTree(b *Btree) {
	if b == nil {
		return
	}

	l := list.New()
	l.PushBack(b)

	for e := l.Front(); e != nil; e = e.Next() {
		left := (e.Value).(*Btree).left
		right := (e.Value).(*Btree).right
		if left != nil {
			l.PushBack(left)
		}

		if right != nil {
			l.PushBack(right)
		}
	}

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(*Btree).value)
	}

}
