//
// 题目：输入一棵二叉树和一个整数， 打印出二叉树中结点值的和为输入整数的所有路径。从树的根结点开始往下一直到叶结点所经过的结点形成一条路径。
//
//  Created by izgnod on 17/10/2.
//  Copyright © 2017年 izgnod. All rights reserved.
//

package tree

import "container/list"
import "fmt"

// Btree 二叉树
type Btree struct {
	value int
	left  *Btree
	right *Btree
}

// FindPath  打印出二叉树中结点值的和为输入整数的所有路径
func FindPath(root *Btree, expectSum int) {
	if root == nil {
		return
	}
	result := list.New()
	findPath(root, 0, expectSum, result)
}

func findPath(root *Btree, curSum, expectSum int, result *list.List) {
	// 前序遍历指定的二叉树，将和放入到result中
	if root != nil {
		curSum += root.value
		result.PushBack(root.value)
		if curSum < expectSum {
			// 当前的和比希望的和小
			findPath(root.left, curSum, expectSum, result)
			findPath(root.right, curSum, expectSum, result)
		} else if curSum == expectSum {
			if root.left == nil && root.right == nil {
				// 正好是叶子结点
				showPath(result)
			}
		}
		result.Remove(result.Back())
	}
}

func showPath(result *list.List) {
	for count, r := 0, result.Front(); r != nil; count, r = count+1, r.Next() {

		if count == 0 {
			fmt.Print(r.Value)
		}
		if count > 0 {
			fmt.Print("->", r.Value)
		}

	}
}
