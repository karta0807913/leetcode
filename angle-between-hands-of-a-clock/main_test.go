package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func angleClock(hour int, minutes int) float64 {
	result := float64((hour%12)*(360/12)) - (float64(minutes) / 60 * 360) + ((float64(minutes) / 60) * (360 / 12))

	if result < 0 {
		result = 360 + result
	}
	if result > 180 {
		result = 360 - result
	}

	return result
}

func TestAngle(t *testing.T) {
	assert.Equal(t, 165.0, angleClock(12, 30), nil)
	assert.Equal(t, 75.0, angleClock(3, 30), nil)
	assert.Equal(t, 7.5, angleClock(3, 15), nil)
	assert.Equal(t, 155.0, angleClock(4, 50), nil)
	assert.Equal(t, 0.0, angleClock(12, 0), nil)
}
