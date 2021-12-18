package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func maxTurbulenceSize(arr []int) (max int) {
	bigger, smaller := 1, 1
	max = 1
	for i := 1; i < len(arr); i++ {
		fmt.Printf("bigger: %d, smaller: %d, currentValue: %d\n", bigger, smaller, arr[i])
		if arr[i-1] < arr[i] {
			bigger += 1
			smaller = 1
		} else if arr[i-1] > arr[i] {
			smaller += 1
			bigger = 1
		} else {
			bigger = 1
			smaller = 1
		}
		if max < bigger {
			max = bigger
		}
		if max < smaller {
			max = smaller
		}
		bigger, smaller = smaller, bigger
	}
	return
}

func TestResult(t *testing.T) {
	assert.Equal(t, 5, maxTurbulenceSize([]int{9, 4, 2, 10, 7, 8, 8, 1, 9}))
	assert.Equal(t, 2, maxTurbulenceSize([]int{4, 8, 12, 16}))
	assert.Equal(t, 1, maxTurbulenceSize([]int{100}))
	assert.Equal(t, 3, maxTurbulenceSize([]int{100, 2, 100}))
	assert.Equal(t, 1, maxTurbulenceSize([]int{100, 100}))
}
