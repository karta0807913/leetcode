package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n > 0 {
		return myPow(x*x, n/2) * myPow(x, n%2)
	}
	return myPow((1/x)*(1/x), n/-2) * myPow(1/x, -n%2)
}

func TestPow(t *testing.T) {
	assert.Equal(t, 9.261000000000001, myPow(2.1, 3), nil)
	assert.Equal(t, 0.25, myPow(2, -2), nil)
	assert.Equal(t, 0.11168554710840535, myPow(8.95371, -1), nil)
	assert.Equal(t, 9.261, myPow(8.84372, -5), nil)
}
