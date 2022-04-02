package main

func countBits(n int) []int {
	dp := make([]int, n+1)
	left, right := 0, 0
	for i := 1; i <= n; i++ {
		if left == right {
			right = i
			left = 0
		}
		dp[i] = dp[left] + 1
		left += 1
	}
	return dp
}
