package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func rangeBitwiseAnd(left int, right int) (ans int) {
	mid := left & right
	for i := 31; i >= 0; i-- {
		ans <<= 1
		mask := 1 << i
		// fmt.Println(ans, mask&mid, mid&^mask, i)
		if mask&mid != 0 && left+mask > right {
			ans |= 1
		}
	}
	return
}

func TestRange(t *testing.T) {
	assert.Equal(t, 4, rangeBitwiseAnd(5, 7))
}
