package main

// TODO: Incorrect :(
import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type DeleteQueue struct {
	char   map[byte]int
	length int
}

func (queue *DeleteQueue) AddChar(c byte) {
	if queue.char == nil || queue.char[c] == 0 {
		queue.char = make(map[byte]int)
		queue.char[c] = 0
	} else {
		tmp := make(map[byte]int)
		tmp[c] = queue.char[c]
		queue.char = tmp
	}
	queue.char[c] += 1
	if queue.char[c] <= 2 || queue.char[c] == 10 || queue.char[c] == 100 {
		queue.length += 1
	}
}

func (queue *DeleteQueue) CompareAndSwap(other *DeleteQueue) {
	if queue.length == other.length {
		for c, n := range other.char {
			if queue.char[c] < n {
				queue.char[c] = n
			}
		}
	} else if queue.length > other.length {
		*queue = *other
	}
}

func getLengthOfOptimalCompression(s string, k int) int {
	dp := make([]DeleteQueue, k+1)
	for i := range s {
		previous := dp[0]
		dp[0].AddChar(s[i])
		for k := 1; k <= (i+1) && k < len(dp); k++ {
			tmp := dp[k]
			dp[k].AddChar(s[i])
			dp[k].CompareAndSwap(&previous)
			// if k == 56 {
			// 	fmt.Printf("dp[k]: %v\n", dp[k])
			// }
			previous = tmp
		}
		fmt.Printf("dp: %v, c: %c\n", dp, s[i])
	}

	minLength := 101
	for _, q := range dp {
		if minLength > q.length {
			minLength = q.length
		}
	}
	return minLength
}

func TestCompression(t *testing.T) {
	assert := require.New(t)
	assert.Equal(
		8,
		getLengthOfOptimalCompression("aabaacdcceefee",
			2,
		),
	)
	{
		q := DeleteQueue{
			char: map[byte]int{
				'a': 5,
				'c': 1,
			},
			length: 12,
		}
		q.AddChar('a')
		assert.Equal(
			DeleteQueue{
				char:   map[byte]int{'a': 6},
				length: 12,
			},
			q,
		)
	}
	assert.Equal(
		12,
		getLengthOfOptimalCompression("ddadccabecddcdbdcaceaceedbebeaadddadcedebbeaedabbcddddecbdaaeeceeebacbacdcccaadcdbaaadcbbabbacea", 56),
	)
	assert.Equal(
		4,
		getLengthOfOptimalCompression("aaabcccd", 2),
	)
	assert.Equal(
		4,
		getLengthOfOptimalCompression("aabaabbcbbbaccc", 6),
	)
	{
		q := DeleteQueue{}
		q.AddChar('a')
		q.AddChar('a')
		q.AddChar('a')
		assert.Equal(
			DeleteQueue{
				char:   map[byte]int{'a': 3},
				length: 2,
			}, q,
		)
	}
	assert.Equal(
		2,
		getLengthOfOptimalCompression("aabbaa", 2),
	)
}
