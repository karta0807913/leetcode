package main

import "math"

func mctFromLeafValues(arr []int) int {
	stack := []int{math.MaxInt}
	result := 0
	for _, num := range arr {
		for stack[len(stack)-1] < num {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			tmp := stack[len(stack)-1]
			if tmp > num {
				tmp = num
			}
			result += mid * tmp
		}
		stack = append(stack, num)
	}
	for len(stack) != 1 {
		result += stack[len(stack)-1] * stack[len(stack)-2]
		stack = stack[:len(stack)-2]
	}
	return result
}
