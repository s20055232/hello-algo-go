package computational_complexity

import (
	"fmt"
	"strconv"
)

type node struct {
	val  int
	next *node
}

func newNode(val int) *node {
	return &node{val: val}
}

func function() int {
	return 0
}

type TreeNode struct {
	Val    any
	Height int
	Left   *TreeNode
	Right  *TreeNode
}

func NewTreeNode(v any) *TreeNode {
	return &TreeNode{
		Val:    v,
		Height: 0,
		Left:   nil,
		Right:  nil,
	}
}

// 常數階，使用的空間跟輸入的資料大小無關
func spaceConstant(n int) {
	// 常數、變數、物件佔用 O(1) 空間
	const a = 0
	b := 0
	nums := make([]int, 1000)
	node := newNode(0)
	// 循環中的“變數”佔用 O(1) 空間
	var c int
	for i := 0; i < n; i++ {
		c = 0
	}
	// 循環中的“函數”佔用 O(1) 空間
	for i := 0; i < n; i++ {
		function()
	}
	b += 0
	c += 0
	nums[0] = 0
	node.val = 0
}

// 線性階
func spaceLinear(n int) {
	// array 與輸入資料長度相關，佔用 O(n) 空間
	_ = make([]int, n)
	// list 與輸入資料長度相關，佔用 O(n) 空間
	var nodes []*node
	for i := 0; i < n; i++ {
		nodes = append(nodes, newNode(i))
	}
	// hashtable 與輸入資料長度相關，佔用 O(n) 空間
	m := make(map[int]string, n)
	for i := 0; i < n; i++ {
		m[i] = strconv.Itoa(i)
	}
}

// 線性階（遞迴實現），callstack 跟輸入資料成線性相關
func spaceLinearRecur(n int) {
	fmt.Println("遞迴 n=", n)
	if n == 1 {
		return
	}
	spaceLinearRecur(n - 1)
}

// 平方階
func spaceQuadratic(n int) {
	// 矩陣佔用 O(n^2) 空間
	numMatrix := make([][]int, n)
	for i := 0; i < n; i++ {
		numMatrix[i] = make([]int, n)
	}
}

// 平方階（遞迴實現），每一輪的 local 變數都會被記錄下來，到最深的時候，使用空間為 O(n^2)
func spaceQuadraticRecur(n int) int {
	if n <= 0 {
		return 0
	}
	nums := make([]int, n)
	fmt.Printf("遞迴 n = %d 中的 nums 長度 = %d \n", n, len(nums))
	return spaceQuadraticRecur(n - 1)
}

// 指數階（建立 full binary tree）O(2^n)
func buildTree(n int) *TreeNode {
	if n == 0 {
		return nil
	}
	root := NewTreeNode(0)
	root.Left = buildTree(n - 1)
	root.Right = buildTree(n - 1)
	return root
}
