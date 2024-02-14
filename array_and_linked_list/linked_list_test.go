package array_and_linkedlist

import (
	"fmt"
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	length := 10
	linkedList := []*ListNode{}
	for i := 0; i < length; i++ {
		linkedList = append(linkedList, NewListNode(i))
	}
	for idx, node := range linkedList {
		if idx != length-1 {
			node.Next = linkedList[idx+1]
		} else {
			node.Next = nil
		}
	}
	curr := linkedList[0]
	for curr.Next != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}
