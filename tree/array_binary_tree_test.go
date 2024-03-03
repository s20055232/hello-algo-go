package tree

import (
	"fmt"
	"testing"
)

func TestArrayBinaryTree(t *testing.T) {
	arr := []any{1, 2, 3, 4, nil, 6, 7, 8, 9, nil, nil, 12, nil, nil, 15}
	root := SliceToTree(arr)
	fmt.Println("Binary Tree initialize")
	fmt.Println("The binary tree presented in an array")
	fmt.Println(arr)
	fmt.Println("The binary tree presented in a linked list")
	PrintTree(root)

	abt := newArrayBinaryTree(arr)
	fmt.Println("current node index is ", 1, "，值为", abt.val(1))
	fmt.Println("left child node index is ", abt.left(1), "，值为", abt.val(abt.left(1)))
	fmt.Println("right child node index is ", abt.right(1), "，值为", abt.val(abt.right(1)))
	fmt.Println("parent node index is ", abt.parent(1), "，值为", abt.val(abt.parent(1)))

	fmt.Println("level order:", abt.levelOrder())
	fmt.Println("in order:", abt.inOrder())
	fmt.Println("pre order:", abt.preOrder())
	fmt.Println("post order:", abt.postOrder())
}
