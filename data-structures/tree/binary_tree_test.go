//
//  normal_tree.go
//
//  Created by izgnod on 17/10/2.
//  Copyright © 2017年 izgnod. All rights reserved.
//

package tree

import "testing"
import "fmt"

func TestTree(t *testing.T) {
	// tree := NewTree()

	// fmt.Printf("%+v\n", tree)

	// tree.Destory()
	// fmt.Printf("%+v\n", tree)

	/*
					A
				B   |    C
			D       | E      F
		G       H   |     I
	*/

	a := NewTree("A")

	b := new(Node)
	b.name = "B"

	c := new(Node)
	c.name = "C"

	d := new(Node)
	d.name = "D"

	e := new(Node)
	e.name = "E"

	f := new(Node)
	f.name = "F"

	g := new(Node)
	g.name = "G"

	h := new(Node)
	h.name = "H"

	i := new(Node)
	i.name = "I"

	// j := new(Node)
	// j.name = "J"

	a.lchild = b
	a.rchild = c

	b.lchild = d

	c.lchild = e
	c.rchild = f

	d.lchild = g
	d.rchild = h

	e.rchild = i
	// e.lchild = j

	fmt.Println("goodsdfj;lsdjf")
	fmt.Printf("%+v\n", a)

	// fmt.Println("PreOrderTraverse:-------------")
	// a.PreOrderTraverse()

	// fmt.Println("InOrderTraverse:--------------")
	// a.InOrderTraverse()

	// fmt.Println("PostOrderTraverse:--------------")
	// a.PostOrderTraverse()

	fmt.Println("LayerOrderTraverse:--------------")
	a.LayerOrderTraverse()
}
