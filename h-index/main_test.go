package main

func hIndex(citations []int) int {
	bulket := make([]int, len(citations)+1)
	for _, cite := range citations {
		if cite >= len(bulket) {
			bulket[len(bulket)-1]++
		} else {
			bulket[cite]++
		}
	}
	result := 0
	sum := 0
	for n := len(bulket) - 1; n >= 0; n-- {
		sum += bulket[n]
		if sum >= n {
			result = n
			break
		}
	}
	return result
}
