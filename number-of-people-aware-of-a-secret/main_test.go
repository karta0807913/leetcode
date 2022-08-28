package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const mod = 1e9 + 7

func peopleAwareOfSecret(n int, delay int, forget int) int {
	currPointer := forget
	date := make([]uint64, forget+n)
	spreading := uint64(0)
	date[currPointer] = 1
	for currPointer < len(date) {
		spreading -= date[currPointer-forget] % mod
		spreading += date[currPointer-delay] % mod
		spreading += mod
		date[currPointer] += (spreading % mod)
		currPointer++
	}
	sCount := uint64(0)
	for i := len(date) - forget; i < len(date); i++ {
		sCount = (sCount + date[i]) % mod
	}
	return int(sCount % mod)
}

func TestSharing(t *testing.T) {
	assert := require.New(t)
	assert.Equal(8, peopleAwareOfSecret(4, 1, 4))
	assert.Equal(5, peopleAwareOfSecret(6, 2, 4))
	assert.Equal(6, peopleAwareOfSecret(4, 1, 3))
	assert.Equal(653668527, peopleAwareOfSecret(684, 18, 496))
}
