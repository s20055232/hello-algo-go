package tree

func preOrder(node *TreeNode) []any {
	slice := []any{}

	var preOrderHelper func(n *TreeNode)
	preOrderHelper = func(n *TreeNode) {
		if n == nil {
			return
		}
		slice = append(slice, n.Val)
		preOrderHelper(n.Left)
		preOrderHelper(n.Right)
	}
	preOrderHelper(node)
	return slice
}

func inOrder(node *TreeNode) []any {
	slice := []any{}

	var preOrderHelper func(n *TreeNode)
	preOrderHelper = func(n *TreeNode) {
		if n == nil {
			return
		}
		preOrderHelper(n.Left)
		slice = append(slice, n.Val)
		preOrderHelper(n.Right)
	}
	preOrderHelper(node)
	return slice
}

func postOrder(node *TreeNode) []any {
	slice := []any{}

	var preOrderHelper func(n *TreeNode)
	preOrderHelper = func(n *TreeNode) {
		if n == nil {
			return
		}
		preOrderHelper(n.Left)
		preOrderHelper(n.Right)
		slice = append(slice, n.Val)
	}
	preOrderHelper(node)
	return slice
}
