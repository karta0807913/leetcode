package main

func trailingZeroes(n int) int {
	result := 0
	for i := 1; i < n; i++ {
		tmp := i
		for tmp%5 == 0 {
			result++
			tmp /= 5
		}
	}
	return result
}
