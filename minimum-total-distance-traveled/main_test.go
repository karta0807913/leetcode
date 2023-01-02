package main

import (
	"math"
	"sort"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type FactoryPair struct {
	PrevFactory []int
	Factory     []int
}

func minimumTotalDistance(robots []int, factory [][]int) int64 {
	var result int64
	sort.Ints(robots)
	sort.Slice(factory, func(i, j int) bool {
		return factory[i][0] < factory[j][0]
	})
	factory = append([][]int{{factory[0][0], 0}}, append(factory, []int{factory[len(factory)-1][0], 0})...)
	leftIndex := 1
	rightIndex := len(factory) - 2
	for len(robots) != 0 && leftIndex < rightIndex {
		for len(robots) != 0 && robots[0] <= factory[leftIndex][0] {
		}

		for len(robots) != 0 && robots[len(robots)-1] >= factory[rightIndex][0] {
		}
	}
	for _, robot := range robots {
		closet := math.MaxInt
		idx := 0
		for i, factory := range factory {
			if factory[1] == 0 {
				continue
			}
			tmp := abs(robot - factory[0])
			if closet > tmp {
				tmp = closet
				idx = i
			}
		}
		factory[idx][1]--
		result += int64(closet)
	}
	return result
}
