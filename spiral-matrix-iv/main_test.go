package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func next(head *ListNode) *ListNode {
	if head.Next == nil {
		head.Val = -1
		return head
	}
	return head.Next
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	c1, c2 := len(matrix[0]), len(matrix)-1
	x, y := -1, 0
	p1, p2 := &x, &y
	d1, d2 := 1, 1
	for c1 != 0 {
		// fmt.Printf("c1: %v, c2: %v\n", c1, c2)
		for i := 0; i != c1; i++ {
			*p1 += d1
			// fmt.Printf("col: %v, row: %v, val: %v\n", y, x, matrix[y][x])
			matrix[y][x] = head.Val
			head = next(head)
		}
		c1, c2, p1, p2, d1, d2 = c2, c1-1, p2, p1, d2, -d1
	}
	return matrix
}
