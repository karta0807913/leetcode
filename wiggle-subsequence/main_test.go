package main

func wiggleMaxLength(nums []int) (ans int) {
	if len(nums) < 2 {
		return len(nums)
	}
	s := 0
	c := nums[0]
	ans = 1
	for idx := 1; idx < len(nums); idx++ {
		if c > nums[idx] && s != 1 {
			s = 1
			ans += 1
		} else if c < nums[idx] && s != -1 {
			s = -1
			ans += 1
		}
		c = nums[idx]
	}
	return ans
}
