package heap

import (
	"testing"

	. "github.com/s20055232/hello-algo-go/pkg"
)

func TestNewMaxHeap(t *testing.T) {
	maxHeap := newMaxHeap1([]int{1, 2, 3, 4, 4, 5, 6})
	PrintHeap(*maxHeap)
	maxHeap2 := newMaxHeap2([]int{1, 2, 3, 4, 4, 5, 6})
	PrintHeap(*maxHeap2)
	maxHeap2 = newMaxHeap2([]int{4, 1, 2, 3, 4, 5, 6})
	PrintHeap(*maxHeap2)
}
