package computational_complexity

// 常數階
func constant(n int) int {
	count := 0
	size := 1000000
	for i := 0; i < size; i++ {
		count++
	}
	return count
}

// 線性階
func linear(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		count++
	}
	return count
}

// 線性階（遍歷 array）
func arrayTraversal(nums []int) int {
	count := 0
	for range nums {
		count++
	}
	return count
}

// 平方階
func quadratic(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			count++
		}
	}
	return count
}

// 平方階（氣泡排序）
func bubbleSort(nums []int) int {
	count := 0
	for i := len(nums) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				tmp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = tmp
				count += 3
			}
		}
	}
	return count
}

// 指數階（循環實現）
// 每一輪 base 都會一分為二，而 count 要根據 base 的大小來添加操作數量
func exponential(n int) int {
	count, base := 0, 1
	for i := 0; i < n; i++ {
		for j := 0; j < base; j++ {
			count++
		}
		base *= 2
	}
	return count
}

// 指數階（遞迴實現）O(2^n)
// 每一輪都會一分為二，最後加上一是因為第一個操作
func expRecur(n int) int {
	if n == 1 {
		return 1
	}
	return expRecur(n-1) + expRecur(n-1) + 1
}

// 對數階（遞迴實現）O(logN)
func logRecur(n float64) int {
	if n <= 1 {
		return 0
	}
	return logRecur(n/2) + 1
}

// 線性對數階 O(n logN)
func linearLogRecur(n float64) int {
	if n <= 1 {
		return 1
	}
	count := linearLogRecur(n/2) + linearLogRecur(n/2)
	for i := 0.0; i < n; i++ {
		count++
	}
	return count
}

// 階乘階（遞迴實現），每一個會分裂出 n 個
func factorialRecur(n int) int {
	if n == 0 {
		return 1
	}
	count := 0
	for i := 0; i < n; i++ {
		count += factorialRecur(n - 1)
	}
	return count
}
