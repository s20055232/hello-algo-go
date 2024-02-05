package computational_complexity

import (
	"testing"
)

type testCase struct {
	num    int
	answer int
}

var testCases []testCase = []testCase{{1, 1}, {2, 3}, {3, 6}, {4, 10}, {5, 15}}

func TestForLoop(t *testing.T) {
	for _, c := range testCases {
		result := forLoop(c.num)
		if result != c.answer {
			t.Errorf("forLoop(%d) = %d; want %d", c.num, result, c.answer)
		}
	}
}

func TestWhileLoop(t *testing.T) {
	for _, c := range testCases {
		result := whileLoop(c.num)
		if result != c.answer {
			t.Errorf("whileLoop(%d) = %d; want %d", c.num, result, c.answer)
		}
	}
}

func TestWhileLoopII(t *testing.T) {
	testCases := []testCase{{1, 1}, {2, 1}, {3, 1}, {4, 5}, {5, 5}}
	for _, c := range testCases {
		result := whileLoopII(c.num)
		if result != c.answer {
			t.Errorf("whileLoopII(%d) = %d; want %d", c.num, result, c.answer)
		}
	}
}

func TestNestedForLoop(t *testing.T) {
	for _, c := range testCases {
		result := nestedForLoop(c.num)
		if result != c.answer {
			t.Errorf("nestedForLoop(%d) = %d; want %d", c.num, result, c.answer)
		}
	}
}
