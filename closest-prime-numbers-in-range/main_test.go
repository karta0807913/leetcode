package main

import (
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

var preProcessedPrimes []int = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409, 419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523, 541, 547, 557, 563, 569, 571, 577, 587, 593, 599, 601, 607, 613, 617, 619, 631, 641, 643, 647, 653, 659, 661, 673, 677, 683, 691, 701, 709, 719, 727, 733, 739, 743, 751, 757, 761, 769, 773, 787, 797, 809, 811, 821, 823, 827, 829, 839, 853, 857, 859, 863, 877, 881, 883, 887, 907, 911, 919, 929, 937, 941, 947, 953, 967, 971, 977, 983, 991, 997}

func sqrt(x int) int {
	j := math.Sqrt(float64(x))
	return int(j) + 1
}

func isPrime(x int) bool {
	if x == 1 {
		return false
	}
	sqrtedN := sqrt(x)
	i := sort.Search(len(preProcessedPrimes), func(i int) bool { return preProcessedPrimes[i] >= sqrtedN })
	for j := 0; j < i; j++ {
		if x%preProcessedPrimes[j] == 0 {
			return false
		}
	}
	return true
}

func closestPrimes(left int, right int) []int {
	if right <= 2 {
		return []int{-1, -1}
	}

	if right <= 1000 {
		i := sort.Search(len(preProcessedPrimes), func(i int) bool { return preProcessedPrimes[i] >= left })

		if i+1 >= len(preProcessedPrimes) {
			return []int{-1, -1}
		}
		minVal := preProcessedPrimes[i+1] - preProcessedPrimes[i]
		ans := []int{preProcessedPrimes[i], preProcessedPrimes[i+1]}
		for j := i + 1; j < len(preProcessedPrimes) && preProcessedPrimes[j] <= right; j++ {
			if preProcessedPrimes[j]-preProcessedPrimes[j-1] < minVal {
				ans = []int{preProcessedPrimes[j-1], preProcessedPrimes[j]}
				minVal = preProcessedPrimes[j] - preProcessedPrimes[j-1]
			}
		}

		if right < ans[1] {
			return []int{-1, -1}
		}
		return ans
	}
	var ans = []int{0, 0}
	prevPrimes := 0
	minGap := 0
	for n := left; n <= right; n++ {
		if isPrime(n) {
			if prevPrimes == 0 {
				prevPrimes = n
			}
			if minGap == 0 || n-prevPrimes < minGap {
				minGap = n - prevPrimes
				ans[0] = prevPrimes
				ans[1] = n
			}
			prevPrimes = n
		}
	}
	if minGap == 0 {
		return []int{-1, -1}
	}
	return ans
}

func TestA(t *testing.T) {
	assert := require.New(t)
	assert.True(isPrime(2))
	assert.True(isPrime(3))
	assert.Equal([]int{1091, 1093}, closestPrimes(1087, 4441))
	assert.Equal([]int{2, 3}, closestPrimes(1, 1001))
	assert.Equal([]int{29, 31}, closestPrimes(19, 31))
	assert.Equal([]int{11, 13}, closestPrimes(10, 19))
	assert.Equal([]int{-1, -1}, closestPrimes(4, 6))
}
