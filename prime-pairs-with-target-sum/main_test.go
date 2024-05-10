package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func isPrime(primes []int, num int) bool {
	if num < 2 {
		return false
	}
	if num == 2 {
		return true
	}
	n := int(math.Sqrt(float64(num)))
	// fmt.Printf("num: %v, n: %v\n", num, n)
	isPrime := true
	for i := 0; i < len(primes) && primes[i] <= n; i++ {
		if num%primes[i] == 0 {
			isPrime = false
			break
		}
	}
	return isPrime
}

func findPrimePairs(n int) [][]int {
	primes := []int{2}
	for num := 3; num <= n/2; num++ {
		if isPrime(primes, num) {
			primes = append(primes, num)
		}
	}
	// fmt.Printf("primes: %v\n", primes)
	var ans [][]int
	for _, prime := range primes {
		if isPrime(primes, n-prime) {
			ans = append(ans, []int{prime, n - prime})
		}
	}
	return ans
}

func TestPrime(t *testing.T) {
	assert := require.New(t)
	assert.Equal([][]int{{3, 7}, {5, 5}}, findPrimePairs(51))
	assert.Equal([][]int{{3, 7}, {5, 5}}, findPrimePairs(10))
	assert.Equal([][]int{{2, 2}}, findPrimePairs(4))
}
