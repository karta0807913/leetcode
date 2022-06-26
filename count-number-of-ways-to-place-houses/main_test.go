package main

func cal(n, houseCount int, hasHouse bool, houses []int) {
	if n == 0 {
		houses[houseCount] += 1
		return
	}
	if hasHouse {
		cal(n-1, houseCount, false, houses)
	} else {
		cal(n-1, houseCount+1, true, houses)
		cal(n-1, houseCount, false, houses)
	}
}

func countHousePlacements(n int) int {
	const mod uint64 = 10e8 + 7
	pprev, prev := uint64(1), uint64(1)
	// if current place has a house, we only can choose pprev count
	// if current place doesn't have a house, we can choose prev count.
	curr := prev + pprev
	for ; n != 1; n-- {
		pprev, prev = prev, curr
		curr = (pprev + prev) % mod
	}
	return int((curr * curr) % mod)
}
