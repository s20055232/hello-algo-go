package array_and_linkedlist

import "testing"

func TestArrayList(t *testing.T) {
	list := newArrayList()
	if list == nil {
		t.Error("create array list failed.")
	}
	if list.size() != 0 {
		t.Error("initial size of array list should be zero.")
	}
	if list.capacity() != 10 {
		t.Error("initial capacity of array list should be ten.")
	}
}
