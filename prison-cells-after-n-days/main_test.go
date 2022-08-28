package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func prisonAfterNDays(cells []int, n int) []int {
	var current int
	for i := range cells {
		current <<= 1
		current |= cells[i]
	}
	cache := make([]int, 256)
	for n != 0 {
		tmp := current
		next := (current << 1) ^ (current >> 1)
		next = (0b11111111 ^ next) & 0b01111110
		if cache[next] != 0 {
			days := (cache[tmp] >> 8) - (cache[next] >> 8) + 1
			n %= days
			current = tmp
			n = (days - n) % days
			for ; n > 0; n-- {
				current = cache[current&(255)]
			}
			break
		}
		n--
		cache[next] = (((cache[tmp] >> 8) + 1) << 8) | tmp
		current = next
	}
	for i := len(cells) - 1; i >= 0; i-- {
		cells[i] = current & 1
		current >>= 1
	}
	return cells
}

func TestDays(t *testing.T) {
	assert := require.New(t)
	assert.Equal(
		[]int{
			0, 0, 1, 0, 0, 1, 0, 0,
		},
		prisonAfterNDays([]int{
			1, 1, 0, 1, 1, 0, 1, 1,
		}, 6),
	)
	assert.Equal(
		[]int{
			0, 0, 1, 1, 0, 0, 0, 0,
		},
		prisonAfterNDays([]int{
			0, 1, 0, 1, 1, 0, 0, 1,
		}, 7),
	)
	assert.Equal(
		[]int{
			0, 0, 1, 1, 1, 1, 1, 0,
		},
		prisonAfterNDays([]int{
			1, 0, 0, 1, 0, 0, 1, 0,
		}, 1000000000),
	)
}
