package iterator

import "errors"

/*
	迭代器模式：又叫游标模式，主要用来遍历容器
*/

type Iterator interface {
	hasNext() bool
	next()
	currentItem() (int, error)
}

type ListIterator struct {
	cursor int
	list   []int // 这里用的list，可以其他
}

func NewListIterator(list []int) *ListIterator {
	return &ListIterator{
		cursor: 0,
		list:   list,
	}
}

func (a *ListIterator) hasNext() bool {
	return a.cursor != len(a.list)
}

func (a *ListIterator) next() {
	a.cursor++
}

func (a *ListIterator) currentItem() (interface{}, error) {
	if a.cursor >= len(a.list) {
		return nil, errors.New("游标超出范围")
	}
	return a.list[a.cursor], nil
}
