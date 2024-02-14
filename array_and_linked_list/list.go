package array_and_linkedlist

type ArrayList struct {
	arrCapacity int
	arr         []int
	arrSize     int
	extendRatio int
}

func newArrayList() *ArrayList {
	return &ArrayList{
		arrCapacity: 10,
		arr:         make([]int, 10),
		arrSize:     0,
		extendRatio: 2,
	}
}

// gets the size of the array list
func (l *ArrayList) size() int {
	return l.arrSize
}

// gets the capacity of the array list
func (l *ArrayList) capacity() int {
	return l.arrCapacity
}

// Any method parameter with 'index' must first use this method for checking
func (l *ArrayList) checkOutOfBound(index int) {
	if index < 0 || index >= l.size() {
		panic("index out of bound")
	}
}

// access element
func (l *ArrayList) get(index int) int {
	l.checkOutOfBound(index)
	return l.arr[index]
}

// update element
func (l *ArrayList) set(num, index int) {
	l.checkOutOfBound(index)
	l.arr[index] = num
}

// add element to the end of the list
func (l *ArrayList) add(num int) {
	if l.arrSize == l.arrCapacity {
		l.extendCapacity()
	}
	l.arr[l.arrSize] = num
	l.arrSize++
}

// insert element to the index
func (l *ArrayList) insert(num, index int) {
	l.checkOutOfBound(index)
	if l.arrSize == l.arrCapacity {
		l.extendCapacity()
	}
	for j := l.arrSize - 1; j >= index; j-- {
		l.arr[j] = l.arr[j-1]
	}
	l.arr[index] = num
	l.arrSize++
}

func (l *ArrayList) extendCapacity() {
	l.arr = append(l.arr, make([]int, l.arrCapacity*(l.extendRatio-1))...)
	l.arrCapacity = len(l.arr)
}

// remove the element
func (l *ArrayList) remove(index int) int {
	l.checkOutOfBound(index)
	num := l.arr[index]
	for j := index; j < l.arrSize-1; j++ {
		l.arr[j] = l.arr[j+1]
	}
	l.arrSize--
	return num
}

func (l *ArrayList) toArray() []int {
	return l.arr[:l.arrSize]
}
