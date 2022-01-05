package main

import "math"

type UnionSet struct {
	islandMap   map[int]int
	islnadCount int
}

func (set *UnionSet) find(x int) int {
	parent, ok := set.islandMap[x]
	if !ok {
		set.islandMap[x] = x
		set.islnadCount += 1
		return x
	}
	if parent == x {
		return x
	}
	set.islandMap[x] = set.find(parent)
	return set.islandMap[x]
}

func (set *UnionSet) union(x, y int) {
	x, y = set.find(x), set.find(y)
	if x != y {
		set.islandMap[y] = x
		set.islnadCount -= 1
	}
}

func removeStones(stones [][]int) int {
	unionSet := UnionSet{
		islandMap:   make(map[int]int),
		islnadCount: 0,
	}
	for _, stone := range stones {
		unionSet.union(stone[0], math.MinInt64+stone[1])
	}
	return len(stones) - unionSet.islnadCount
}
