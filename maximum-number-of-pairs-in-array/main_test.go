package main

func numberOfPairs(nums []int) []int {
	m := make(map[int]int)
	for _, val := range nums {
		m[val] += 1
	}
	ans := make([]int, 2)
	for _, v := range m {
		ans[0] += v / 2
		ans[1] += v % 2
	}
	return ans
}
