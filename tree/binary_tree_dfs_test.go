package tree

import (
	"fmt"
	"testing"
)

func TestPreInPostOrderTraveral(t *testing.T) {
	root := SliceToTree([]any{1, 2, 3, 4, 5, 6, 7})
	fmt.Println("Tree initialize: ")
	PrintTree(root)

	nums := preOrder(root)
	fmt.Println("printed nums from running preorder: ", nums)

	nums = inOrder(root)
	fmt.Println("printed nums from running inorder: ", nums)

	nums = postOrder(root)
	fmt.Println("printed nums from running postorder: ", nums)
}
