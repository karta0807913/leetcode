package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	result := ""
	basicOffset := numRows + numRows - 2
	for headOffset := 0; headOffset < numRows; headOffset++ {
		offsetA := basicOffset - headOffset*2
		offsetB := basicOffset - offsetA
		for idx := headOffset; idx < len(s); {
			if offsetA != 0 {
				result += string(s[idx])
				idx += offsetA
			}

			if idx >= len(s) {
				break
			}

			if offsetB != 0 {
				result += string(s[idx])
				idx += offsetB
			}
		}
	}
	return result
}

func TestConvert(t *testing.T) {
	assert.Equal(t, "PAHNAPLSIIGYIR", convert("PAYPALISHIRING", 3), nil)
	assert.Equal(t, "PINALSIGYAHRPI", convert("PAYPALISHIRING", 4), nil)
	assert.Equal(t, "PHASIYIRPLIGAN", convert("PAYPALISHIRING", 5), nil)
}
