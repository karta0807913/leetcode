package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

const mod = 1e9 + 7

func recursive(index int, nums []int, mask int, cache []map[int]int) int {
	mask |= (1 << index)
	if ans, ok := cache[index][mask]; ok {
		return ans
	}
	fmt.Printf("nums[index]: %2v, mask: %14b, %d\n", nums[index], mask, mask)
	ans := 0
	for i := 0; i < len(nums); i++ {
		if mask&(1<<i) != 0 {
			continue
		}
		if nums[index]%nums[i] == 0 || nums[i]%nums[index] == 0 {
			ans += recursive(i, nums, mask, cache)
			ans %= mod
		}
	}
	cache[index][mask] = ans
	return ans
}

func specialPerm(nums []int) int {
	cache := make([]map[int]int, len(nums))
	for i := range cache {
		cache[i] = make(map[int]int)
		cache[i][(1<<len(nums))-1] = 1
	}
	ans := 0
	for i := range nums {
		fmt.Println("START", ans)
		ans += recursive(i, nums, 0, cache)
		ans %= mod
		fmt.Printf("cache: %v\n", cache)
	}
	return ans
}

func TestPerm(t *testing.T) {
	assert := require.New(t)
	assert.Equal(
		2,
		specialPerm([]int{1, 4, 3}),
	)
	assert.Equal(
		2,
		specialPerm([]int{2, 3, 6}),
	)
	assert.Equal(
		178290591,
		specialPerm([]int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192}),
	)
}
