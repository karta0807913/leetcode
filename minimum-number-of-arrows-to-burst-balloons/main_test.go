package main

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	prev := points[0]
	count := 1

	for idx := 1; idx < len(points); idx++ {
		if points[idx][0] > prev[1] {
			prev = points[idx]
			count += 1
		}
	}
	return count
}
