package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func recuriveDevide(divided, divisor int64) (int64, int) {
	if divided == 0 {
		return 0, 0
	}
	times := 0
	if divided >= divisor+divisor {
		divided, times = recuriveDevide(divided, divisor+divisor)
		times += times
		if divided == 0 {
			return 0, times
		}
	}
	for divided >= divisor {
		divided -= divisor
		times += 1
	}
	return divided, times
}

func divide(dividend int, divisor int) int {
	nagtive := false
	if dividend < 0 {
		nagtive = !nagtive
		dividend = -dividend
	}
	if divisor < 0 {
		nagtive = !nagtive
		divisor = -divisor
	}
	_, times := recuriveDevide(int64(dividend), int64(divisor))
	if nagtive {
		times = -times
	}
	if times > 2147483647 || times < -2147483648 {
		return 2147483647
	}
	return times
}

func TestDivide(t *testing.T) {
	assert.Equal(t, 7/3, divide(7, 3), nil)
	assert.Equal(t, 7/-3, divide(7, -3), nil)
	assert.Equal(t, 0/1, divide(0, 1), nil)
	assert.Equal(t, 1/1, divide(1, 1), nil)
	assert.Equal(t, 2147483647, divide(-2147483648, -1), nil)
}
