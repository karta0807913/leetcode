package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func decodeDigits(s string) (string, int) {
	result := 0
	for '0' <= s[0] && s[0] <= '9' {
		result *= 10
		result += int(s[0] - '0')
		s = s[1:]
	}
	return s, result
}

func extractString(s string) (extracted string, remain string) {
	stringBuilder := strings.Builder{}
	for len(s) != 0 {
		if s[0] == ']' {
			s = s[1:]
			break
		}
		if '0' <= s[0] && s[0] <= '9' {
			var repeat int
			s, repeat = decodeDigits(s)
			extracted, remain = extractString(s[1:])
			for i := 0; i < repeat; i++ {
				stringBuilder.WriteString(extracted)
			}
			s = remain
		} else if 'a' <= s[0] && s[0] <= 'z' {
			stringBuilder.WriteByte(s[0])
			s = s[1:]
		}
	}
	return stringBuilder.String(), s
}

func decodeString(s string) string {
	s, _ = extractString(s)
	return s
}

func TestDecodeString(t *testing.T) {
	assert := require.New(t)
	assert.Equal("a", decodeString("1[a]"))
	assert.Equal("aaa", decodeString("3[a]"))
	assert.Equal("aaabaa", decodeString("3[a]b2[a]"))
	assert.Equal("aaabaaaaabaa", decodeString("2[3[a]b2[a]]"))
}
