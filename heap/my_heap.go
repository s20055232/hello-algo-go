package heap

import "fmt"

type myMaxHeap []any

func (h *myMaxHeap) Left(i int) int {
	return 2*i + 1
}

func (h *myMaxHeap) Right(i int) int {
	return 2*i + 2
}

func (h *myMaxHeap) Parent(i int) int {
	return (i - 1) / 2
}

func (h myMaxHeap) Peek() any {
	return h[0]
}

func (h myMaxHeap) size() int {
	return len(h)
}

func (h *myMaxHeap) Push(val any) {
	*h = append(*h, val)
	h.SiftUp(h.size() - 1)
}

func (h *myMaxHeap) SiftUp(i int) {
	for {
		p := h.Parent(i)
		if p < 0 || (*h)[i].(int) <= (*h)[p].(int) {
			break
		}
		h.Swap(i, p)
		i = p
	}
}

func (h *myMaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myMaxHeap) Pop() any {
	if h.size() == 0 {
		fmt.Print("error")
		return nil
	}
	h.Swap(0, h.size()-1)
	val := (*h)[h.size()-1]
	*h = (*h)[:h.size()-1]
	h.SiftDown(0)
	return val
}

func (h *myMaxHeap) SiftDown(i int) {
	for {
		l, r, max := h.Left(i), h.Right(i), i
		if l < h.size() && (*h)[l].(int) > (*h)[i].(int) {
			max = l
		}
		if r < h.size() && (*h)[r].(int) > (*h)[i].(int) {
			max = r
		}
		if max == i {
			break
		}
		h.Swap(i, max)
		i = max
	}
}
