package tree

import (
	"fmt"
	"testing"
)

func SliceToTreeDFS(arr []any, i int) *TreeNode {
	if i < 0 || i >= len(arr) || arr[i] == nil {
		return nil
	}
	root := NewTreeNode(arr[i])
	root.Left = SliceToTreeDFS(arr, 2*i+1)
	root.Right = SliceToTreeDFS(arr, 2*i+2)
	return root
}

// SliceToTree 将切片反序列化为二叉树
func SliceToTree(arr []any) *TreeNode {
	return SliceToTreeDFS(arr, 0)
}
func TestLevelOrder(t *testing.T) {
	root := SliceToTree([]any{1, 2, 3, 4, 5, 6, 7})
	fmt.Println("Tree initialize: ")
	PrintTree(root)
	nums := levelOrder(root)
	fmt.Println("printed nums from level order: ", nums)
}
