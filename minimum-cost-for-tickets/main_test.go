package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func mincostTickets(days []int, costs []int) int {
	cache := make([]int, days[len(days)-1]+1)

	dayIter := 0
	for day := 1; day < len(cache); day++ {
		oneDayCost := cache[max(0, day-1)] + costs[0]
		sevenDaysCost := cache[max(0, day-7)] + costs[1]
		thirtyDaysCost := cache[max(0, day-30)] + costs[2]
		cache[day] = min(min(oneDayCost, sevenDaysCost), thirtyDaysCost)
		if day != days[dayIter] {
			cache[day] = min(cache[day-1], cache[day])
		} else {
			dayIter += 1
		}
		// fmt.Printf("cache: %v\n", cache)
	}
	return cache[len(cache)-1]
}

func TestMinCost(t *testing.T) {
	assert.Equal(t, 11, mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}))
	assert.Equal(t, 17, mincostTickets([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15}))
}
