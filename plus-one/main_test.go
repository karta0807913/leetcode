package main

func plusOne(digits []int) []int {
	carry := 1
	for idx := len(digits) - 1; idx >= 0 && carry != 0; idx-- {
		cache := digits[idx] + carry
		digits[idx] = cache % 10
		carry = cache / 10
	}
	if carry == 1 {
		return append([]int{1}, digits...)
	}
	return digits
}
