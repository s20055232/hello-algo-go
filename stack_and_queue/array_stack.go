package stack_and_queue

type arrayStack struct {
	data []int
}

func newArrayStack() *arrayStack {
	return &arrayStack{
		data: make([]int, 0, 16),
	}
}

func (s *arrayStack) size() int {
	return len(s.data)
}

func (s *arrayStack) isEmpty() bool {
	return len(s.data) == 0
}

func (s *arrayStack) push(v int) {
	s.data = append(s.data, v)
}

func (s *arrayStack) pop() any {
	result := s.peek()
	s.data = s.data[:len(s.data)-1]
	return result
}

func (s *arrayStack) peek() any {
	if s.isEmpty() {
		return nil
	}
	return s.data[len(s.data)-1]
}

func (s *arrayStack) toSlice() []int {
	return s.data
}
