package heap

type intHeap struct {
	nums    []any
	maxHeap bool
}

func (h *intHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	h.nums = append(h.nums, x.(int))
}

func (h *intHeap) Pop() any {
	last := h.nums[h.Len()-1]
	h.nums = h.nums[:h.Len()-1]
	return last
}

func (h intHeap) Len() int {
	return len(h.nums)
}
func (h intHeap) Less(i, j int) bool {
	if h.maxHeap {
		return h.nums[i].(int) > h.nums[j].(int)
	}
	return h.nums[i].(int) < h.nums[j].(int)
}
func (h intHeap) Swap(i, j int) {
	h.nums[i], h.nums[j] = h.nums[j], h.nums[i]
}

func (h intHeap) Top() any {
	return h.nums[0]
}
