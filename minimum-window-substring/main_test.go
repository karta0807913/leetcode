package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func minWindow(s string, t string) string {
	wants := make(map[byte]int)
	current := make(map[byte]int)
	for i := range t {
		wants[t[i]] += 1
	}
	var ans []int
	reachRequest := 0
	var left, right int
	for right < len(s) {
		// search right to reach the requirement
		for ; right < len(s); right++ {
			b := s[right]
			_, ok := wants[b]
			if ok {
				current[b]++
				if current[b] == wants[b] {
					reachRequest++
					if reachRequest == len(wants) {
						right++
						break
					}
				}
			}
		}
		if reachRequest != len(wants) {
			break
		}
		// fmt.Printf("s[left:right]: %v\n", s[left:right])
		// shrink left to get minimal range
		for ; left < right; left++ {
			b := s[left]
			_, ok := wants[b]
			if ok {
				current[b] -= 1
				if current[b] == wants[b]-1 {
					if ans == nil || right-left < ans[1]-ans[0] {
						// fmt.Printf("HIs[left:right]: %v\n", s[left:right])
						ans = []int{left, right}
					}
					reachRequest -= 1
					left++
					break
				}
			}
		}
		// fmt.Printf("s[left:right]: %v\n", s[left:right])
	}
	// fmt.Printf("ans: %v\n", ans)
	if ans == nil {
		return ""
	}
	return s[ans[0]:ans[1]]
}

func TestWindow(t *testing.T) {
	assert.Equal(t, "", minWindow("a", "ab"))
	assert.Equal(t, "BANC", minWindow("ADOBECODEBANC", "ABC"))
	assert.Equal(t, "a", minWindow("a", "a"))
	assert.Equal(t, "a", minWindow("ab", "a"))
	assert.Equal(t, "", minWindow("a", "aa"))
}
