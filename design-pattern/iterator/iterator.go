package iterator

import "errors"

/*
	迭代器模式：又叫游标模式，主要用来遍历容器
	一个完整的迭代器模式，一般会涉及容器和容器迭代器两部分内容。
		为了达到基于接口而非实现编程的目的，容器又包含容器接口、容器实现类.
		迭代器又包含迭代器接口、迭代器实现类。
		容器中需要定义 iterator() 方法，用来创建迭代器。
		迭代器接口中需要定义 hasNext()、currentItem()、next() 三个最基本的方法。
        容器对象通过依赖注入传递到迭代器类中。
*/
// 容器接口
type Container interface {
	getIterator() Iterator
	indexItem(idx int) int
	size() int
}

// 迭代类接口
type Iterator interface {
	hasNext() bool
	next()
	currentItem() (int, error)
}

// 容器的实现类
type ConcreteContainer struct {
	list []int // 这里用的list，可以其他
}

func NewConcreteContainer(list []int) *ConcreteContainer {
	return &ConcreteContainer{list: list}
}

func (l *ConcreteContainer) size() int {
	return len(l.list)
}

func (l *ConcreteContainer) indexItem(idx int) int {
	return l.list[idx]
}

func (l *ConcreteContainer) getIterator() Iterator {
	return NewContainerIterator(l)

}

// 迭代类
type ContainerIterator struct {
	list   Container
	cursor int
}

func NewContainerIterator(list Container) *ContainerIterator {
	return &ContainerIterator{
		cursor: 0,
		list:   list,
	}
}

func (l *ContainerIterator) hasNext() bool {
	return l.cursor != l.list.size()
}

func (l *ContainerIterator) next() {
	l.cursor++
}

func (l *ContainerIterator) currentItem() (int, error) {
	if l.cursor >= l.list.size() {
		return 0, errors.New("游标超出范围")
	}
	return l.list.indexItem(l.cursor), nil
}
