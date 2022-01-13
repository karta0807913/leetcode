package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func left(b []byte) int {
	if len(b) == 0 {
		return 0
	}
	c := 1
	prev := b[len(b)-1]
	for idx := len(b) - 2; idx >= 0; idx-- {
		if prev != b[idx] {
			break
		}
		c += 1
	}
	return c
}

func right(s string, k int) int {
	if len(s) == 0 {
		return 0
	}
	c := 1
	prev := s[0]
	for idx := 1; idx < len(s); idx++ {
		if prev != s[idx] || c == k {
			break
		}
		c += 1
	}
	return c
}

func removeDuplicates(s string, k int) string {
	answer := make([]byte, 0, len(s))
	for len(s) != 0 {
		r := right(s, k)
		if r == k {
			s = s[r:]
			continue
		}

		// fmt.Println(string(answer), s)
		l := left(answer)
		if l+r >= k && answer[len(answer)-1] == s[0] {
			r = k - l
			s = s[r:]
			answer = answer[:len(answer)-l]
			continue
		}
		answer = append(answer, s[:r]...)
		s = s[r:]
	}
	return string(answer)
}

func TestDuplicates(t *testing.T) {
	assert.Equal(t, 3, left([]byte("bbaaa")))
	assert.Equal(t, 1, left([]byte("a")))
	assert.Equal(t, 3, right("aaaa", 3))
	assert.Equal(t, 1, right("a", 1))

	assert.Equal(t, "ps", removeDuplicates("pbbcggttciiippooaais", 2))
	assert.Equal(t, "ab", removeDuplicates("abccbb", 2))
	assert.Equal(t, "aa", removeDuplicates("deeedbbcccbdaa", 3))
	assert.Equal(t, "aa", removeDuplicates("aa", 3))
	assert.Equal(t, "", removeDuplicates("aaa", 3))
	assert.Equal(t, "", removeDuplicates("baaabb", 3))
	assert.Equal(t, "aabb", removeDuplicates("aabb", 3))
}
