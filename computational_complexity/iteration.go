package computational_complexity

func forLoop(n int) int {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	return res
}

func whileLoop(n int) int {
	res := 0
	i := 1
	for i <= n {
		res += i
		i++
	}
	return res
}

func whileLoopII(n int) int {
	res := 0
	i := 1
	for i <= n {
		res += i
		i++
		i *= 2
	}
	return res
}

func nestedForLoop(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == j {
				res += i
			}
		}
	}
	return res
}
