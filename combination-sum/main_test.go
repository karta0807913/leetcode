package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func combinationSum(candidates []int, target int) (ans [][]int) {
	ans = [][]int{}
	if target == 0 {
		return [][]int{{}}
	}
	for idx, candidate := range candidates {
		if candidate > target {
			continue
		}
		for _, subAns := range combinationSum(candidates[idx:], target-candidate) {
			ans = append(ans, append(subAns, candidate))
		}
	}
	return ans
}

// slower version
// func combinationSum(candidates []int, target int) [][]int {
// 	cache := make([][][]int, target+1)
// 	for i := 0; i < len(candidates); i++ {
// 		if candidates[i] > target {
// 			continue
// 		}
// 		cache[candidates[i]] = append(cache[candidates[i]], []int{candidates[i]})
// 		for n := candidates[i] + 1; n <= target; n++ {
// 			for _, val := range cache[n-candidates[i]] {
// 				cache[n] = append(cache[n], append(append([]int{}, val...), candidates[i]))
// 			}
// 		}
// 	}
// 	// fmt.Printf("cache: %v\n", cache)
// 	return cache[target]
// }

func TestSum(t *testing.T) {
	assert.ElementsMatch(t, [][]int{
		{2, 2, 2, 2}, {2, 3, 3}, {3, 5},
	}, combinationSum([]int{2, 3, 5}, 8))
	assert.ElementsMatch(t, [][]int{
		{2, 2, 3}, {7},
	}, combinationSum([]int{2, 3, 6, 7}, 7))

	assert.ElementsMatch(t, [][]int{}, combinationSum([]int{2}, 1))

	assert.ElementsMatch(t, [][]int{
		{2, 2, 2, 2, 2, 3}, {2, 2, 2, 2, 5}, {2, 2, 3, 3, 3}, {2, 3, 3, 5}, {3, 5, 5},
	}, combinationSum([]int{2, 3, 5}, 13))
}
