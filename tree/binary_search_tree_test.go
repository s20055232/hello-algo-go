package tree

import (
	"fmt"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	bst := newBinaryTree()
	nums := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15}
	for _, num := range nums {
		bst.insert(num)
	}
	fmt.Println("binary tree initialize:")
	bst.Print()

	node := bst.search(7)
	fmt.Println("The node that was found is ", node, ", node value is = ", node.Val)

	bst.insert(16)
	fmt.Println("binary tree after a new node was inserted:")
	bst.Print()

	bst.remove(1)
	fmt.Println("binary tree after the node 1 was deleted:")
	bst.Print()

	bst.remove(2)
	fmt.Println("binary tree after the node 2 was deleted:")
	bst.Print()

	bst.remove(4)
	fmt.Println("binary tree after the node 4 was deleted:")
	bst.Print()
}
