package main

func minSteps(n int) int {
	step := 0
	cb := 0
	currentCount := 1
	for currentCount < n {
		if n%currentCount == 0 {
			step += 2
			cb = currentCount
			currentCount += currentCount
		} else {
			currentCount += cb
			step += 1
		}
	}
	return step
}
