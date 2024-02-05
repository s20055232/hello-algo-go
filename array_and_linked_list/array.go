package array_and_linkedlist

import "math/rand"

// array 隨機訪問，O(1)
func randomAccess(nums []int) int {
	randomIndex := rand.Intn(len(nums))
	return nums[randomIndex]
}

// 擴容，要找到更大的連續空間，並將元素搬移過去，O(n)
func extend(nums []int, enlarge int) []int {
	res := make([]int, len(nums)+enlarge)
	for i, num := range nums {
		res[i] = num
	}
	return res
}

// 插入，要將插入的索引之後的元素都往後移動一位，O(n)
func insert(nums []int, num int, index int) {
	for i := len(nums) - 1; i > index; i-- {
		nums[i] = nums[i-1]
	}
	nums[index] = num
}

// 刪除，要將刪除的元素後面的元素都往前移動一位，O(n)
func remove(nums []int, index int) {
	for i := index; i < len(nums)-1; i++ {
		nums[i] = nums[i+1]
	}
}

func traverse(nums []int) {
	count := 0
	// 透過索引遍歷 array
	for i := 0; i < len(nums); i++ {
		count += nums[i]
	}
	count = 0
	// 直接遍歷 array
	for _, num := range nums {
		count += num
	}
	count = 0
	// 同時遍歷 array 的索引跟數值
	for i, num := range nums {
		count += num
		count += nums[i]
	}
}

// 線性查找，O(n)
func find(nums []int, target int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == target {
			return i
		}
	}
	return -1
}
