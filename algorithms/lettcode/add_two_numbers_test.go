package lettcode

import (
	"fmt"
	"testing"
)

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	listNode := &ListNode{}
	borrow := false
	for ret := listNode; l1 != nil || l2 != nil || borrow; {
		var tmp int
		if borrow {
			tmp++
			borrow = false
		}
		if l1 != nil && l2 != nil {
			tmp += l1.Val + l2.Val
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 != nil && l2 == nil {
			tmp += l1.Val
			l1 = l1.Next
		} else if l1 == nil && l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}
		if tmp >= 10 {
			borrow = true
			tmp = tmp % 10
		}
		(*ret).Val = tmp

		if l1 != nil || l2 != nil || borrow {
			(*ret).Next = &ListNode{}
		} else {
			break
		}
		ret = ret.Next
	}
	return listNode
}

func TestAddTwoNumbers(t *testing.T) {
	tmp1 := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}
	tmp2 := ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}}}
	v2 := addTwoNumbers(&tmp1, &tmp2)
	for {
		fmt.Printf("val:%d\n", v2.Val)
		if v2.Next == nil {
			break
		}
		v2 = v2.Next
	}
}
