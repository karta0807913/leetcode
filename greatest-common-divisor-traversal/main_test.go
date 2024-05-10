package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func canTraverseAllPairs(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	var primeList []int = []int{2, 3, 5, 7, 11}
	isPrime := func(num int) bool {
		for _, val := range primeList {
			if val*val > num {
				return true
			}
			if num%val == 0 {
				return false
			}
		}
		return true
	}
	groupList := func(num int) []int {
		var ans []int
		lastPrime := primeList[len(primeList)-1] + 2
		// generate prime
		for ; lastPrime*lastPrime < num; lastPrime += 2 {
			if isPrime(lastPrime) {
				primeList = append(primeList, lastPrime)
			}
		}

		for _, val := range primeList {
			if num == 1 {
				break
			}
			if num%val == 0 {
				ans = append(ans, val)
				for num%val == 0 {
					num /= val
				}
			}
		}
		if num != 1 {
			ans = append(ans, num)
		}
		return ans
	}

	groupFind := make(map[int]int)
	groupCount := 0
	var findParent func(int) int
	findParent = func(children int) int {
		parent, ok := groupFind[children]
		if !ok {
			groupFind[children] = children
			groupCount++
			return children
		}
		if parent != children {
			groupFind[children] = findParent(parent)
		}
		return groupFind[children]
	}
	setParent := func(parent, children int) {
		parent, children = findParent(parent), findParent(children)
		if parent != children {
			groupCount--
		}
		groupFind[children] = parent
	}

	for _, num := range nums {
		if num == 1 {
			return false
		}
		// fmt.Printf("groupList(%d): %v\n", num, groupList(num))
		groups := groupList(num)
		parent, groups := groups[0], groups[1:]
		findParent(parent) // init parent
		for _, val := range groups {
			setParent(parent, val)
		}
	}
	// fmt.Printf("groupCount: %v\n", groupCount)
	// fmt.Printf("groupFind: %+v\n", groupFind)

	return groupCount == 1
}

func TestCanTraverseAllPairs(t *testing.T) {
	assert := require.New(t)
	assert.False(canTraverseAllPairs([]int{42, 40, 45, 42, 50, 33, 30, 45, 33, 45, 30, 36, 44, 1, 21, 10, 40, 42, 42}))
	assert.True(canTraverseAllPairs([]int{2, 3, 6}))
}
