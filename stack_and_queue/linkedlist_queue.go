package stack_and_queue

import "container/list"

type linkedListQueue struct {
	data *list.List
}

func newLinkedListQueue() *linkedListQueue {
	return &linkedListQueue{
		data: list.New(),
	}
}

func (s *linkedListQueue) push(value int) {
	s.data.PushBack(value)
}

func (s *linkedListQueue) pop() any {
	if s.isEmpty() {
		return nil
	}
	e := s.data.Front()
	s.data.Remove(e)
	return e.Value
}

func (s *linkedListQueue) isEmpty() bool {
	return s.data.Len() == 0
}

func (s *linkedListQueue) peek() any {
	if s.isEmpty() {
		return nil
	}
	return s.data.Front().Value
}

func (s *linkedListQueue) size() int {
	return s.data.Len()
}

func (s *linkedListQueue) toList() *list.List {
	return s.data
}
