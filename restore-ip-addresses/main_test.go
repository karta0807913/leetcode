package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func removeLeadingZero(s string) string {
	for idx := range s {
		if s[idx] != '0' {
			return s[idx:]
		}
	}
	return "0"
}

func isLeadingZeros(s string) bool {
	if len(s) == 1 {
		return false
	}
	isZero := false
	for idx := range s {
		if s[idx] == '0' {
			isZero = true
		} else {
			if isZero {
				return true
			} else {
				return false
			}
		}
	}
	return isZero
}

func resursiveAddress(prefix, current string, count int, result []string) []string {
	if count == 4 && len(current) != 0 {
		return result
	}
	if count == 4 && len(current) == 0 {
		result = append(result, prefix[1:])
		return result
	}
	for idx := 1; idx <= 3 && idx <= len(current); idx++ {
		if isLeadingZeros(current[:idx]) {
			continue
		}
		i, _ := strconv.Atoi(current[:idx])
		if i <= 255 {
			result = resursiveAddress(prefix+"."+current[:idx], current[idx:], count+1, result)
		}
	}
	return result
}

func restoreIpAddresses(s string) []string {
	return resursiveAddress("", s, 0, make([]string, 0))
}

func TestAddress(t *testing.T) {
	// assert.Equal(t, []string{"255.255.11.135", "255.255.111.35"}, restoreIpAddresses("25525511135"), nil)
	assert.Equal(t, []string{"0.0.0.0"}, restoreIpAddresses("0000"), nil)
	assert.Equal(t, []string{"1.1.1.1"}, restoreIpAddresses("1111"), nil)
	assert.Equal(t, []string{"0.10.0.10", "0.100.1.0"}, restoreIpAddresses("010010"), nil)
	assert.Equal(t, []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"}, restoreIpAddresses("101023"), nil)
	// assert.Equal(t, "12", removeLeadingZero("00012"), nil)
	// assert.Equal(t, "0", removeLeadingZero("000000"), nil)
	// assert.Equal(t, "1", removeLeadingZero("0000001"), nil)
	// assert.Equal(t, "100000001", removeLeadingZero("000100000001"), nil)
	// assert.Equal(t, true, isLeadingZeros("00"), nil)
	// assert.Equal(t, false, isLeadingZeros("10"), nil)
}
