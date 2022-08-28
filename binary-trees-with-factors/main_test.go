package main

import "sort"

const modulo = 1e9 + 7

func numFactoredBinaryTrees(arr []int) int {
	dp := make(map[int]int, len(arr))
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	for _, val := range arr {
		dp[val] = 1
	}
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i]%arr[j] == 0 {
				dp[arr[i]] += (dp[arr[j]] * dp[arr[i]/arr[j]]) % modulo
			}
		}
	}
	total := 0
	for _, val := range dp {
		total += val
		total %= modulo
	}
	return total
}
