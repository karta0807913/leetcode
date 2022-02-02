package main

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for idx := 1; idx < len(dp); idx++ {
		for _, num := range nums {
			if num <= idx {
				dp[num] += dp[idx-num]
			}
		}
	}
	return dp[target]
}
