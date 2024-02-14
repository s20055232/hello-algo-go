package stack_and_queue

import (
	"testing"
)

func TestLinkedListStack(t *testing.T) {
	l := newLinkedListStack()
	l.push(1)
	l.push(2)
	l.push(3)
	l.push(4)
	v := l.pop()
	if 4 != v {
		t.Errorf("Expected 4, got: %v", v)
	}
	if l.isEmpty() {
		t.Error("Expected stack is not empty")
	}
	if 3 != l.peek() {
		t.Error("Expected stack's top is 3")
	}
	if l.size() != 3 {
		t.Error("Expected stack's size is 3")
	}
	l.pop()
	l.pop()
	l.pop()
	v = l.pop()
	if v != nil {
		t.Error("Expected stack is empty")
	}
	if !l.isEmpty() {
		t.Error("Expected stack is empty")
	}
}
