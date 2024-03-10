package heap

import (
	"container/heap"
	"fmt"
	"testing"

	. "github.com/s20055232/hello-algo-go/pkg"
)

func TestIntHeap(t *testing.T) {
	h := &intHeap{
		nums:    []any{2, 1, 5},
		maxHeap: false,
	}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", h.nums[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}

	h = &intHeap{
		nums:    []any{2, 1, 5},
		maxHeap: true,
	}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("\nmaximum: %d\n", h.nums[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	size := h.Len()
	fmt.Println("\nsize of heap is", size)

	isEmpty := size == 0
	fmt.Println("size is empty:", isEmpty)
}

func TestMyHeap(t *testing.T) {
	maxHeap := myMaxHeap([]any{9, 8, 6, 6, 7, 5, 2, 1, 4, 3, 6, 2})
	fmt.Printf("Input the array and initialize the heap\n")
	PrintHeap(maxHeap)

	peek := maxHeap.Peek()
	fmt.Println("top element of the heap is:", peek)

	maxHeap.Push(7)
	fmt.Println("after element: 7 push into heap")
	PrintHeap(maxHeap)

	peek = maxHeap.Pop()
	fmt.Printf("after element: %d pop from heap\n", peek)
	PrintHeap(maxHeap)

	size := maxHeap.size()
	fmt.Println("\nsize of heap is", size)

	isEmpty := size == 0
	fmt.Println("size is empty:", isEmpty)
}
