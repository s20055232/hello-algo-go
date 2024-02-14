package array_and_linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

// 在 head node 後面插入節點
func inertNode(head *ListNode, node *ListNode) {
	n1 := head.Next
	head.Next = node
	node.Next = n1
}

// 刪除節點
func removeItem(node *ListNode) {
	if node.Next == nil {
		return
	}
	node.Next = node.Next.Next
}

// 訪問索引為 index 的節點
func access(head *ListNode, index int) *ListNode {
	for i := 0; i < index; i++ {
		if head == nil {
			return nil
		}
		head = head.Next
	}
	return head
}

// 查找值為 target 的首個節點
func findNode(head *ListNode, target int) int {
	index := 0
	for head != nil {
		if head.Val == target {
			return index
		}
		index++
		head = head.Next
	}
	return -1
}

type DoublyListNode struct {
	Val  int
	Next *DoublyListNode
	Prev *DoublyListNode
}

func NewDoublyListNode(val int) *DoublyListNode {
	return &DoublyListNode{
		Val:  val,
		Next: nil,
		Prev: nil,
	}
}
