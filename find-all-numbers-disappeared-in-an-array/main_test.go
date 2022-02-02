package main

func GO(nums []int, t int) {
	for nums[t-1] != t {
		nums[t-1], t = t, nums[t-1]
	}
}

func findDisappearedNumbers(nums []int) (ans []int) {
	for _, n := range nums {
		GO(nums, n)
	}
	for i, n := range nums {
		if n-1 != i {
			ans = append(ans, n)
		}
	}
	return
}
