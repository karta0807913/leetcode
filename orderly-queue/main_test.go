package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func orderlyQueue(s string, k int) string {
	if k == 1 {
		buf := make([]byte, len(s), len(s)*2)
		copy(buf, s)
		min := s
		for i := range s {
			buf = append(buf, s[i])
			if tmp := string(buf[i : len(s)+i]); tmp < min {
				min = tmp
			}
		}
		return min
	}
	buf := []byte(s)
	sort.Slice(buf, func(i, j int) bool {
		return buf[i] < buf[j]
	})
	return string(buf)
}

func TestQueue(t *testing.T) {
	assert := require.New(t)
	assert.Equal("isxrx", orderlyQueue("xisxr", 1))
}
