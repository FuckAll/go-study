//
//  normal_tree.go
//
//  Created by izgnod on 17/10/2.
//  Copyright © 2017年 izgnod. All rights reserved.
//

package tree

import (
	"fmt"
)

/**
* 实现一个简单的树结构
方法:
   1. NewTree
*/

type Node struct {
	name   string
	lchild *Node
	rchild *Node
	custom interface{}
}

// NewTree return tree's root
func NewTree(root string) *Node {
	t := new(Node)
	t.name = root
	return t
}

// DestoryTree tree
func (n *Node) DestoryTree() {
	n.name = ""
	n.lchild = nil
	n.rchild = nil
	n.custom = nil
}

// // CreateTree create tree by nodes
// func (n *Node) CreateTree(nodes []*Node) {

// }

// PreOrderTraverse 前序遍历
func (n *Node) PreOrderTraverse() {

	if n == nil {
		return
	}
	fmt.Println(n.name)

	n.lchild.PreOrderTraverse()
	n.rchild.PreOrderTraverse()
}

// InOrderTraverse 中序遍历
func (n *Node) InOrderTraverse() {
	if n == nil {
		return
	}

	n.lchild.InOrderTraverse()

	fmt.Println(n.name)

	n.rchild.InOrderTraverse()

}

// PostOrderTraverse 后序遍历
func (n *Node) PostOrderTraverse() {

	if n == nil {
		return
	}

	n.lchild.PostOrderTraverse()
	n.rchild.PostOrderTraverse()
	fmt.Println(n.name)
}

// LayerOrderTraverse 按层遍历
// 将根节点先放入队列中，然后从队列中以此遍历，如果发现结点有左右孩子，放在队列之中
func (n *Node) LayerOrderTraverse() {
	tmp := []*Node{}

	if n == nil {
		fmt.Println("tree empty.")
	}
	tmp = append(tmp, n)

	for i := 0; ; i++ {
		if n.lchild != nil {
			tmp = append(tmp, n.lchild)
		}

		if n.rchild != nil {
			tmp = append(tmp, n.rchild)
		}

		if i+1 >= len(tmp) {
			break
		}
		n = tmp[i+1]
	}

	for _, t := range tmp {
		fmt.Println(t.name)
	}
}
