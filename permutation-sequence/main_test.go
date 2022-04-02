package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func recursive(n []byte, k, r int) {
	if len(n) == 1 {
		return
	}
	r /= len(n)
	// fmt.Printf("r: %v\n", r)
	i := k / r
	k %= r
	t := n[i]
	for idx := i; idx > 0; idx-- {
		n[idx] = n[idx-1]
	}
	n[0] = t
	// fmt.Printf("n: %v\n", string(n))
	recursive(n[1:], k, r)
}

func getPermutation(n int, k int) string {
	r := 1
	b := make([]byte, n)
	for i := range b {
		r *= i + 1
		b[i] = '1' + byte(i)
	}
	recursive(b, k-1, r)
	return string(b)
}

func TestPermutation(t *testing.T) {
	assert.Equal(t, "213", getPermutation(3, 3))
	assert.Equal(t, "2314", getPermutation(4, 9))
}
