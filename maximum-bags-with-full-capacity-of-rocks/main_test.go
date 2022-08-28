package main

import "sort"

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	sorted := make([]int, len(capacity))
	for i := range capacity {
		sorted[i] = rocks[i] - capacity[i]
	}
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] > sorted[j] })
	for i := range sorted {
		additionalRocks += sorted[i]
		if additionalRocks < 0 {
			return i
		}
	}
	return len(capacity)
}
