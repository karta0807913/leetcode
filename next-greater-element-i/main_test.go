package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	cache := make(map[int]int)
	stack := make([]int, 0, len(nums2))
	for idx := len(nums2) - 1; idx >= 0; idx-- {
		cache[nums2[idx]] = idx
		target := -1
		for len(stack) != 0 {
			n := len(stack) - 1
			pop := stack[n]
			stack = stack[:n]
			if pop > nums2[idx] {
				stack = append(stack, pop)
				target = pop
				break
			}
		}
		stack = append(stack, nums2[idx])
		nums2[idx] = target
	}
	for idx := range nums1 {
		nums1[idx] = nums2[cache[nums1[idx]]]
	}
	return nums1
}
