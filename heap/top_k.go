package heap

import "container/heap"

func topKHeap(nums []int, k int) *intHeap {
	// 把東西交給我，我來幫你擴充跟運行
	h := &intHeap{}
	heap.Init(h)
	for i := 0; i < k; i++ {
		heap.Push(h, nums[i])
	}
	for i := k; i < len(nums); i++ {
		if nums[i] > h.Top().(int) {
			heap.Pop(h)
			heap.Push(h, nums[i])
		}
	}
	return h
}
