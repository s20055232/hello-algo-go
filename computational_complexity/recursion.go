package computational_complexity

import "container/list"

func recursion(n int) int {
	if n == 1 {
		return 1
	}
	res := recursion(n - 1)
	return res + n
}

func tailRecur(n int, res int) int {
	if n == 0 {
		return res
	}
	return tailRecur(n-1, res+n)
}

func fib(n int) int {
	if n == 1 || n == 0 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func forLoopRecur(n int) int {
	// 建立雙向鏈結串列
	stack := list.New()
	res := 0
	for i := n; i > 0; i-- {
		stack.PushBack(i)
	}
	for stack.Len() != 0 {
		res += stack.Back().Value.(int)
		stack.Remove(stack.Back())
	}
	return res
}
