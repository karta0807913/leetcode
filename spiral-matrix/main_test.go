package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func spiralOrder(matrix [][]int) []int {
	c1, c2 := len(matrix[0]), len(matrix)-1
	ans := make([]int, 0, len(matrix)*len(matrix[0]))
	x, y, d1, d2 := -1, 0, 1, 1
	p1, p2 := &x, &y
	for c1 != 0 {
		// fmt.Printf("c1: %v, c2: %v\n", c1, c2)
		for i := 0; i != c1; i++ {
			*p1 += d1
			// fmt.Printf("col: %v, row: %v, val: %v\n", y, x, matrix[y][x])
			ans = append(ans, matrix[y][x])
		}
		c1, c2, p1, p2, d1, d2 = c2, c1-1, p2, p1, d2, -d1
		// c1, c2 = c2, c1-1
		// p1, p2 = p2, p1
		// d1, d2 = d2, -d1
	}
	return ans
}

func TestOrder(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}, spiralOrder([][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}))
	assert.Equal(t, []int{1, 2, 3, 4, 5}, spiralOrder([][]int{
		{1}, {2}, {3}, {4}, {5},
	}))
	assert.Equal(t, []int{1, 2, 3, 4, 5}, spiralOrder([][]int{
		{1, 2, 3, 4, 5},
	}))
}
