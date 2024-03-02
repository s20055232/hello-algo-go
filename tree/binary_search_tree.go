package tree

import (
	"fmt"
)

func PrintTree(root *TreeNode) {
	printTreeHelper(root, nil, false)
}
func printTreeHelper(root *TreeNode, prev *trunk, isRight bool) {
	if root == nil {
		return
	}
	prevStr := "    "
	trunk := newTrunk(prev, prevStr)
	printTreeHelper(root.Right, trunk, true)
	if prev == nil {
		trunk.str = "———"
	} else if isRight {
		trunk.str = "/———"
		prevStr = "   |"
	} else {
		trunk.str = "\\———"
		prev.str = prevStr
	}
	showTrunk(trunk)
	fmt.Println(root.Val)
	if prev != nil {
		prev.str = prevStr
	}
	trunk.str = "   |"
	printTreeHelper(root.Left, trunk, false)
}

type trunk struct {
	prev *trunk
	str  string
}

func newTrunk(prev *trunk, str string) *trunk {
	return &trunk{
		prev: prev,
		str:  str,
	}
}

func showTrunk(t *trunk) {
	if t == nil {
		return
	}

	showTrunk(t.prev)
	fmt.Print(t.str)
}

type TreeNode struct {
	Val    any
	Height int
	Left   *TreeNode
	Right  *TreeNode
}

type binarySearchTree struct {
	root *TreeNode
}

func newBinaryTree() *binarySearchTree {
	bst := &binarySearchTree{}
	bst.root = nil
	return bst
}

func newTreeNode(v any) *TreeNode {
	return &TreeNode{
		Val: v, Height: 0, Left: nil, Right: nil,
	}
}

func (bst *binarySearchTree) getRoot() *TreeNode {
	return bst.root
}

func (bst *binarySearchTree) search(num int) *TreeNode {
	node := bst.root
	for node != nil {
		if node.Val.(int) < num {
			node = node.Left
		} else if node.Val.(int) > num {
			node = node.Right
		} else {
			break
		}
	}
	return node
}

func (bst *binarySearchTree) insert(num int) {
	cur := bst.root
	newNode := newTreeNode(num)
	if cur == nil {
		bst.root = newNode
		return
	}

	var pre *TreeNode = nil

	for cur != nil {
		if cur.Val == nil {
			// if encounter same value, do nothing and return
			return
		}
		pre = cur
		if cur.Val.(int) > num {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}

	if num > pre.Val.(int) {
		pre.Right = newNode
	} else {
		pre.Left = newNode
	}
}

func (bst *binarySearchTree) remove(num int) {
	cur := bst.root
	var pre *TreeNode = nil

	for cur != nil {
		if cur.Val.(int) == num {
			break
		}
		pre = cur
		if cur.Val.(int) < num {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}
	fmt.Println("previous node is:", pre)
	fmt.Println("current node is:", cur)
	if cur == nil {
		return
	}
	fmt.Println("previous node is:", pre)
	fmt.Println("current node is:", cur)
	if cur.Left == nil || cur.Right == nil {
		// use child to catch child node
		var child *TreeNode
		if cur.Left != nil {
			child = cur.Left
		} else {
			child = cur.Right
		}
		if cur != bst.root {
			if pre.Left == cur {
				pre.Left = child
			} else {
				pre.Right = child
			}
		} else {
			bst.root = child
		}
	} else {
		tmp := cur.Right
		for tmp.Left != nil {
			tmp = tmp.Left
		}
		bst.remove(tmp.Val.(int))
		cur.Val = tmp.Val
	}
}

func (bst *binarySearchTree) print() {
	PrintTree(bst.root)
}
