package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func robotWithString(s string) string {
	p, t := make([]byte, 0, len(s)), make([]byte, 0, len(s))
	stack := make([]int, len(s))
	stack[len(stack)-1] = len(stack) - 1
	for i := len(s) - 2; i >= 0; i-- {
		stack[i] = i
		if s[stack[i+1]] < s[i] {
			stack[i] = stack[i+1]
		}
	}
	for idx := 0; idx < len(s); {
		if len(t) == 0 || s[stack[idx]] < t[len(t)-1] {
			t = append(t, s[idx])
			idx += 1
		} else {
			p = append(p, t[len(t)-1])
			t = t[:len(t)-1]
		}
	}
	for idx := len(t) - 1; idx >= 0; idx-- {
		p = append(p, t[idx])
	}
	return string(p)
}

func TestString(t *testing.T) {
	assert := require.New(t)
	assert.Equal(
		"fnohopzv",
		robotWithString("vzhofnpo"),
	)
}
