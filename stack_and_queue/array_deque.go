package stack_and_queue

type arrayDeque struct {
	nums        []int
	front       int
	queSize     int
	queCapacity int
}

func newArrayDeque(capacity int) *arrayDeque {
	return &arrayDeque{
		nums:        make([]int, 0, capacity),
		front:       0,
		queSize:     0,
		queCapacity: capacity,
	}
}

func (q *arrayDeque) size() int {
	return q.queSize
}

func (q *arrayDeque) isEmpty() bool {
	return q.queSize == 0
}

// 接受任意數字並得到一個合法的索引
//
// 在處理負數時可以從頭回看（逆著看，跟 Python 一樣），超出索引也不會出錯
func (q *arrayDeque) index(i int) int {
	return (i + q.queCapacity) % q.queCapacity
}

func (q *arrayDeque) pushFirst(num int) {
	if q.queSize == q.queCapacity {
		return
	}
	q.front = q.index(q.front - 1)
	q.nums[q.front] = num
	q.queSize++
}

func (q *arrayDeque) pushLast(num int) {
	if q.queSize == q.queCapacity {
		return
	}
	rear := q.front + q.queSize
	rear = q.index(rear)
	q.nums[rear] = num
	q.queSize++
}

func (q *arrayDeque) peekFirst() any {
	if q.isEmpty() {
		return nil
	}
	return q.nums[q.front]
}

func (q *arrayDeque) popFirst() any {
	result := q.peekFirst()
	q.front = q.index(q.front + 1)
	q.queSize--
	return result
}

func (q *arrayDeque) peekLast() any {
	if q.isEmpty() {
		return nil
	}
	last := q.index(q.front + q.queSize - 1)
	return q.nums[last]
}

func (q *arrayDeque) popLast() any {
	result := q.peekLast()
	q.queSize--
	return result
}

func (q *arrayDeque) toSlice() []int {
	res := make([]int, q.queSize)
	for i, j := 0, q.front; i < q.queSize; i++ {
		res[i] = q.nums[q.index(j)]
		j++
	}
	return res
}
