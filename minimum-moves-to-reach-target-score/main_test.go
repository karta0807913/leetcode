package main

func minMoves(target int, maxDoubles int) int {
	if target == 1 {
		return 0
	}
	n := 0
	for ; target != 1 && maxDoubles != 0; n++ {
		n += target & 1
		maxDoubles--
		target >>= 1
	}
	return n + target - 1
}
