package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type SubStr struct {
	ExistsMap    int64
	CurrentCount int
	Weight       int
}

func (subStr *SubStr) AddChar(b byte) {
	var n int64 = 1 << ((b - 'A') * 2)
	// if char not exists
	if subStr.ExistsMap&n == 0 {
		subStr.ExistsMap |= n
		subStr.CurrentCount += 1
		return
	}

	// if char appear only once
	if subStr.ExistsMap&(n<<1) == 0 {
		subStr.ExistsMap |= n << 1
		subStr.CurrentCount -= 1
	}
}

func uniqueLetterString(s string) int {
	dp := make([]SubStr, len(s)+1)
	subStr := make([]*SubStr, 0)
	result := 0
	for i := 0; i < len(s); i++ {
		currentChar := s[i]
		dp[i].CurrentCount = 0
		dp[i].Weight = 1
		subStr = append(subStr, &dp[i])
		fmt.Printf("currentChar: %c\n", currentChar)
		leftPointer := 0
		sameMap := make(map[int64]*SubStr)
		for rightPointer, subSet := range subStr {
			subSet.AddChar(currentChar)
			result += subSet.CurrentCount * subSet.Weight
			subStr[leftPointer], subStr[rightPointer] =
				subStr[rightPointer], subStr[leftPointer]
			leftPointer += 1
			fmt.Printf("subSet: %v\n", subSet)
			fmt.Printf("result: %v\n", result)
			if s, ok := sameMap[subSet.ExistsMap]; ok {
				s.Weight += subSet.Weight
				leftPointer -= 1
			} else {
				sameMap[subSet.ExistsMap] = subSet
			}
		}
		subStr = subStr[:leftPointer]
	}
	return result
}

func TestUnique(t *testing.T) {
	assert := require.New(t)
	assert.Equal(6, uniqueLetterString("AAB"))
	assert.Equal(12, uniqueLetterString("LEET"))
	assert.Equal(2, uniqueLetterString("AA"))
	assert.Equal(10, uniqueLetterString("ABC"))
	assert.Equal(92, uniqueLetterString("LEETCODE"))
	assert.Equal(1989, uniqueLetterString("LEEEEEEEEEEEEEEEEEEEEEEEEEEEEEETTTTTCODEEEEEEEEEEEEEE"))
}
