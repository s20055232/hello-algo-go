package stack_and_queue

type arrayQueue struct {
	nums        []int
	front       int
	queSize     int
	queCapacity int
}

func newArrayQueue(queCapacity int) *arrayQueue {
	return &arrayQueue{
		nums:        make([]int, 0, queCapacity),
		front:       0,
		queSize:     0,
		queCapacity: queCapacity,
	}
}

func (q *arrayQueue) size() int {
	return q.queSize
}

func (q *arrayQueue) isEmpty() bool {
	return q.queSize == 0
}

func (q *arrayQueue) push(num int) {
	if q.queSize == q.queCapacity {
		return
	}
	rear := (q.front + q.queSize) % q.queCapacity
	q.nums[rear] = num
	q.queSize++
}

func (q *arrayQueue) pop() any {
	num := q.peek()
	q.front = (q.front + 1) % q.queCapacity
	q.queSize--
	return num
}

func (q *arrayQueue) peek() any {
	if q.isEmpty() {
		return nil
	}
	return q.nums[q.front]
}

func (q *arrayQueue) toSlice() []int {
	rear := (q.front + q.queSize)
	return q.nums[q.front:rear]
}
