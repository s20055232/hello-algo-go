package stack_and_queue

import "container/list"

type linkedListStack struct {
	data *list.List
}

func newLinkedListStack() *linkedListStack {
	return &linkedListStack{
		data: list.New(),
	}
}

func (s *linkedListStack) push(value int) {
	s.data.PushBack(value)
}

func (s *linkedListStack) pop() any {
	if s.isEmpty() {
		return nil
	}
	e := s.data.Back()
	s.data.Remove(e)
	return e.Value
}

func (s *linkedListStack) peek() any {
	if s.isEmpty() {
		return nil
	}
	return s.data.Back().Value
}

func (s *linkedListStack) size() int {
	return s.data.Len()
}

func (s *linkedListStack) isEmpty() bool {
	return s.data.Len() == 0
}

func (s *linkedListStack) toList() *list.List {
	return s.data
}
