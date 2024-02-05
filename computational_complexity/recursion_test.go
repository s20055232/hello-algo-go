package computational_complexity

import "testing"

func TestRecursion(t *testing.T) {
	for _, c := range testCases {
		result := recursion(c.num)
		if result != c.answer {
			t.Errorf("recursion(%d) = %d; want %d", c.num, result, c.answer)
		}
	}
}

func TestTailRecursion(t *testing.T) {
	for _, c := range testCases {
		result := tailRecur(c.num, 0)
		if result != c.answer {
			t.Errorf("tailRecursion(%d) = %d; want %d", c.num, result, c.answer)
		}
	}
}

func TestFib(t *testing.T) {
	testCases := []testCase{{1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {7, 13}}
	for _, c := range testCases {
		result := fib(c.num)
		if result != c.answer {
			t.Errorf("fib(%d) = %d; want %d", c.num, result, c.answer)
		}
	}
}

func TestForLoopRecur(t *testing.T) {
	for _, c := range testCases {
		result := forLoopRecur(c.num)
		if result != c.answer {
			t.Errorf("forLoopRecur(%d) = %d; want %d", c.num, result, c.answer)
		}
	}
}
