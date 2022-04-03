package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func numSquares(n int) int {
	current := []int{n}
	next := []int{}
	i := 1
LOOP:
	for len(current) != 0 {
		for _, n := range current {
			begin := int(math.Sqrt(float64(n)))
			for ps := begin; ps >= begin/2; ps-- {
				n := n - ps*ps
				if n == 0 {
					break LOOP
				}
				next = append(next, n)
			}
		}
		current, next = next, current
		next = next[:0]
		i += 1
	}
	return i
}

func TestSquares(t *testing.T) {
	assert.Equal(t, 2, numSquares(13))
	assert.Equal(t, 3, numSquares(12))
}
