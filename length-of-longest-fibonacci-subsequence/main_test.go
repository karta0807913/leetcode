package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type Pair struct {
	First, Second int
}

func (pair Pair) Key() string {
	return fmt.Sprintf("%d:%d", pair.First, pair.Second)
}

func lenLongestFibSubseq(arr []int) int {
	existsMap := make(map[int]bool, len(arr))
	dp := make(map[string]int)
	result := 0
	existsMap[arr[len(arr)-1]] = true
	existsMap[arr[len(arr)-2]] = true
	for i := len(arr) - 3; i >= 0; i-- {
		existsMap[arr[i]] = true
		for j := len(arr) - 1; j > i; j-- {
			if existsMap[arr[i]+arr[j]] {
				n := 2
				pair := Pair{
					First:  arr[i],
					Second: arr[j],
				}
				key := pair.Key()
				// fmt.Printf("pair: %v\n", pair)
				for existsMap[pair.First+pair.Second] {
					pair.First, pair.Second = pair.Second, pair.First+pair.Second
					// fmt.Printf("pair.Key(): %v\n", pair.Key())
					// fmt.Printf("n: %v\n", n)
					k := pair.Key()
					if dp[k] != 0 {
						n += dp[k] - 1
						break
					}
					n += 1
				}
				dp[key] = n
				if n > result {
					result = n
				}
				// fmt.Printf("dp: %v\n", dp)
			}
		}
	}
	return result
}

func TestSubseq(t *testing.T) {
	assert := require.New(t)
	assert.Equal(5, lenLongestFibSubseq([]int{
		1, 2, 3, 4, 5, 6, 7, 8,
	}))
	assert.Equal(3, lenLongestFibSubseq([]int{
		1, 3, 7, 11, 12, 14, 18,
	}))
}
