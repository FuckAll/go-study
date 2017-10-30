package list

import (
	"errors"
	"fmt"
)

// Node 列表节点结构体
type Node struct {
	Value interface{}
	Pre   *Node
	Next  *Node
}

// List 列表结构体
type List struct {
	Length int
	Head   *Node
	Tail   *Node
}

// NewNode 如何创造一个新的节点
func NewNode(value interface{}) *Node {
	n := new(Node)
	n.Value = value
	return n
}

// Prepend 头部插入Node方法
func (l *List) Prepend(n *Node) {
	if l.Length == 0 {
		l.Head = n
		l.Tail = n
	} else {
		l.Head.Pre = n
		n.Next = l.Head
		l.Head = n
	}

	l.Length++
}

// Append 将Node追加到最后
func (l *List) Append(n *Node) {
	if l.Length == 0 {
		l.Head = n
		l.Tail = n
	} else {
		l.Tail.Next = n
		n.Pre = l.Tail
		l.Tail = n
	}
	l.Length++
}

// Insert 指定位置插入
/*
	Index 从0开始
*/
func (l *List) Insert(index int, n interface{}) error {
	// 超过长度和空的List都无法插入
	if index > l.Length {
		return errors.New("Index out of range")
	}

	if index == 0 || l.Length == 0 {
		l.Prepend(NewNode(n))
		return nil
	}

	if l.Length == index {
		l.Append(NewNode(n))
		return nil
	}

	indexNode, err := l.Get(index)
	if err != nil {
		return err
	}

	newNode := NewNode(n)

	newNode.Pre = indexNode.Pre
	newNode.Next = indexNode

	indexNode.Pre.Next = newNode
	indexNode.Pre = newNode
	l.Length++

	return nil
}

// Remove 删除指定位置的Node
func (l *List) Remove(index int) error {
	if l.Length < 1 {
		return errors.New("List is empty")
	}

	if index > l.Length-1 {
		return errors.New("Index out of range")
	}

	if index == 0 {
		if l.Length == 1 {
			l.Head = nil
			l.Tail = nil
		} else {
			l.Head = l.Head.Next
			l.Head.Pre = nil
		}
		l.Length--
		return nil
	}

	if index == l.Length-1 {
		if l.Length == 1 {
			l.Head = nil
			l.Tail = nil
		} else {
			l.Tail = l.Tail.Pre
			l.Tail.Next = nil
		}

		l.Length--
		return nil
	}

	node, err := l.Get(index)
	if err != nil {
		return err
	}

	node.Next.Pre = node.Pre
	node.Pre.Next = node.Next
	l.Length--
	return nil

}

// Show List
func (l *List) Show() {
	for n := l.Head; n != nil; n = n.Next {
		fmt.Printf("%+v\n", n.Value)
	}
}

// Get By Index
func (l *List) Get(index int) (*Node, error) {
	if index > l.Length {
		return nil, errors.New("Index out of range")
	}

	n := l.Head
	for i := 0; i < index; i++ {
		n = n.Next
	}

	return n, nil
}

// Find 查找某个Node对应的Index
func (l *List) Find(n *Node) (int, error) {
	// 两个数据结构要一样,并且值相等
	index := 0
	for node := l.Head; node != nil; node = node.Next {
		if n.Value == node.Value {
			return index, nil
		}
		index++
	}
	return -1, errors.New("Item not found")
}

// Contact 合并链表
func (l *List) Contact(nextList *List) {
	if nextList.Length != 0 {
		if l.Length == 0 {
			l = nextList
		} else {
			l.Tail.Next = nextList.Head
			nextList.Head.Pre = l.Tail
			l.Length = l.Length + nextList.Length
		}
	}
}
