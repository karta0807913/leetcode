package main

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func compare(version1, version2 string) int {
	v1, _ := strconv.Atoi(version1)
	v2, _ := strconv.Atoi(version2)
	if v1 < v2 {
		return -1
	} else if v1 > v2 {
		return 1
	}
	return 0
}

func removeFootZeros(v []string) []string {
	for i := len(v) - 1; i >= 0; i-- {
		vi, _ := strconv.Atoi(v[i])
		if vi == 0 {
			v = v[:len(v)-1]
		} else {
			break
		}
	}
	return v
}

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	v1 = removeFootZeros(v1)
	v2 = removeFootZeros(v2)
	length := 0
	if len(v1) < len(v2) {
		length = len(v1)
	} else {
		length = len(v2)
	}
	for i := 0; i < length; i++ {
		result := compare(v1[i], v2[i])
		if result != 0 {
			return result
		}
	}
	if len(v1) < len(v2) {
		return -1
	} else if len(v1) > len(v2) {
		return 1
	}
	return 0
}

func TestCompare(t *testing.T) {
	// assert.Equal(t, -1, compareVersion("0.1", "1.1"), nil)
	// assert.Equal(t, 1, compareVersion("1.0.1", "1"), nil)
	// assert.Equal(t, -1, compareVersion("7.5.2.4", "7.5.3"), nil)
	assert.Equal(t, 0, compareVersion("1.0", "1.0.0"), nil)
}
