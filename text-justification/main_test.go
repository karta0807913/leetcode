package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeSpaces(n int) string {
	s := ""
	for ; n > 0; n-- {
		s += " "
	}
	return s
}

func makeStr(buffer []string, maxWidth, strLen int) (res string) {
	if len(buffer) == 1 {
		return buffer[0] + makeSpaces(maxWidth-len(buffer[0]))
	}
	n := len(buffer) - 1
	spaces := (maxWidth - strLen + (n + 1))
	extraSpaces := spaces % n
	s := spaces / n
	padding := makeSpaces(s)
	for idx := 0; idx < n; idx++ {
		res += buffer[idx] + padding
		if extraSpaces > 0 {
			res += " "
		}
		extraSpaces -= 1
	}
	res += buffer[n]
	return
}

func fullJustify(words []string, maxWidth int) (ans []string) {
	var bufferLen, offset int
	for offset != len(words) {
		if bufferLen+len(words[offset]) > maxWidth {
			ans = append(ans, makeStr(words[:offset], maxWidth, bufferLen))
			words = words[offset:]
			bufferLen = 0
			offset = 0
		}
		bufferLen += len(words[offset]) + 1
		offset += 1
	}
	if len(words) != 0 {
		s := ""
		for _, w := range words {
			s += w + " "
		}
		if len(s) > maxWidth {
			s = s[:maxWidth]
		} else {
			s += makeSpaces(maxWidth - len(s))
		}
		ans = append(ans, s)
	}
	return
}

func TestMake(t *testing.T) {
	assert.Equal(t, "a  b", makeStr([]string{"a", "b"}, 4, 4))
	assert.Equal(t, "a b c", makeStr([]string{"a", "b", "c"}, 5, 6))
	assert.Equal(t, "a  b  c", makeStr([]string{"a", "b", "c"}, 7, 6))
	assert.Equal(t, []string{"a  b  c"}, fullJustify([]string{"a", "b", "c"}, 7))
}
