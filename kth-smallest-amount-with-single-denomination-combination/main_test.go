package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func GCD(a int64, b int64) int64 {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

type SDC [][]int64

func CreateSDC(coins []int) SDC {
	sdc := make([][]int64, len(coins))
	for i, coin := range coins {
		coin := int64(coin)
		prevCoins := []int64{1}
		for i, currentCoins := range sdc[:i+1] {
			tmp := currentCoins
			for _, prevCoin := range prevCoins {
				next := (coin * prevCoin) / GCD(prevCoin, coin)
				currentCoins = append(currentCoins, next)
			}
			sdc[i] = currentCoins
			prevCoins = tmp
		}
	}
	for i := range sdc {
		sort.Slice(sdc[i], func(a, b int) bool {
			return sdc[i][a] < sdc[i][b]
		})
	}
	return sdc
}

func (sdc SDC) GetNumberBeforeN(n int64) int64 {
	total := int64(0)
	for i, coins := range sdc {
		currentTotal := int64(0)
		for _, coin := range coins {
			if coin > n {
				break
			}
			currentTotal += n / int64(coin)
		}
		if i%2 != 0 {
			currentTotal = -currentTotal
		}
		total += currentTotal
	}
	// fmt.Printf("n: %v, total: %v\n", n, total)
	return total
}

func findKthSmallest(coins []int, _k int) int64 {
	sort.Ints(coins)
	k := int64(_k)

	upperBound := int64(coins[len(coins)-1])*k + 1
	lowerBound := int64(1)
	sdc := CreateSDC(coins)
	// fmt.Printf("sdc: %v\n", sdc)
	answer := int64(0)
	for lowerBound < upperBound {
		mid := (upperBound-lowerBound)/2 + lowerBound
		if count := sdc.GetNumberBeforeN(mid); count < k {
			lowerBound = mid + 1
		} else if count >= k {
			answer = mid
			upperBound = mid
		}
		// fmt.Printf("lowerBound: %v, upperBound: %v\n", lowerBound, upperBound)
	}
	return answer
}

func TestFindKthSmallest(t *testing.T) {
	assert := require.New(t)
	assert.Equal(int64(20368), findKthSmallest([]int{2, 25, 17, 19, 3, 13, 15, 14, 1, 4}, 20368))
	assert.Equal(int64(12), findKthSmallest([]int{10, 3, 7, 5}, 7))
	assert.Equal(int64(36), findKthSmallest([]int{10, 6}, 8))
	assert.Equal(int64(9), findKthSmallest([]int{3, 6, 9}, 3))
	assert.Equal(int64(12), findKthSmallest([]int{5, 2}, 7))
}
