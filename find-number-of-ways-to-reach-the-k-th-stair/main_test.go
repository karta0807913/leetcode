package main

func waysToReachStair(k int) int {
	ans := 0
	for i := 0; i < 31; i++ {
		upper := 1 << (i)
		lower := upper - (i + 1)
		if lower <= k && k <= upper {
			operationCount := upper - k
			totalSlot := i + 1
			// C(totalSlot, operationCount)
			res := uint64(1)
			for i := 0; i < operationCount; i++ {
				res = res * uint64(totalSlot-i) / uint64(i+1)
			}
			// fmt.Println(i, totalSlot, operationCount)
			ans += int(res)
		}
	}
	return ans
}
