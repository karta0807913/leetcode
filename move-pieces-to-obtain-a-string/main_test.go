package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// assert.False(canChange("R__L", "_LR_"))
func canChange(start string, target string) bool {
	count := 0
	for i := range start {
		if target[i] == 'L' {
			count += 1
		}
		if start[i] == 'L' {
			count--
			if count < 0 {
				return false
			}
		}
		if start[i] == 'R' || target[i] == 'R' {
			if count != 0 {
				return false
			}
		}
	}

	for i := len(start) - 1; i >= 0; i-- {
		if target[i] == 'R' {
			count += 1
		}
		if start[i] == 'R' {
			count--
			if count < 0 {
				return false
			}
		}
		if start[i] == 'L' || target[i] == 'L' {
			if count != 0 {
				return false
			}
		}
	}
	return count == 0
}

// L_
// _L

// __L__L__R___L___
// _____L___L_R_L____

func TestChange(t *testing.T) {
	assert := require.New(t)
	assert.False(canChange("R__L", "_LR_"))
	assert.False(canChange("L_", "_L"))
	assert.True(canChange("R_", "_R"))
	assert.False(canChange("_R", "R_"))
	assert.False(canChange("__L__L__R___L___", "_____L___L_R_L____"))
	assert.True(canChange("_L__R__R_", "L______RR"))
	assert.False(canChange("R_L_", "__LR"))
	assert.False(canChange("_______RL___", "____RL______"))
	assert.False(canChange("____", "R_L_"))
	assert.False(canChange("L___", "R_L_"))
}
