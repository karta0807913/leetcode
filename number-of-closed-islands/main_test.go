package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func min(a int, b ...int) int {
	if len(b) == 0 {
		return a
	}

	val := min(b[0], b[0:]...)
	if a < val {
		return a
	}
	return val
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func recordConvert(landIdx, targetIdx int, landConvertMap map[int]int) int {
	l := landIdx
	h := targetIdx
	if targetIdx < landIdx {
		l, h = h, l
	}
	v, ok := landConvertMap[h]
	if ok && v != l {
		l = recordConvert(v, l, landConvertMap)
	}
	landConvertMap[h] = l
	return l
}

func checkSpace(landIdx, targetIdx int, landConvertMap map[int]int) int {
	if targetIdx >= 2 {
		if landIdx != -1 && landIdx != targetIdx {
			recordConvert(landIdx, targetIdx, landConvertMap)
		}
		return targetIdx
	}
	return landIdx
}

func checkLandStatus(colIdx, rowIdx int, grid [][]int, landConvertMap map[int]int) (int, bool) {
	isEdage := false
	landIdx := -1
	if colIdx == 0 {
		isEdage = true
	} else {
		landIdx = checkSpace(landIdx, grid[colIdx-1][rowIdx], landConvertMap)
	}

	if rowIdx == 0 {
		isEdage = true
	} else {
		landIdx = checkSpace(landIdx, grid[colIdx][rowIdx-1], landConvertMap)
	}

	isEdage = isEdage || rowIdx+1 == len(grid[0]) || colIdx+1 == len(grid)

	return landIdx, isEdage
}

func closedIsland(grid [][]int) int {
	isNotIslandMap := make(map[int]bool)
	landConvertMap := make(map[int]int)
	currentIsland := 2
	for colIdx := range grid {
		for rowIdx := range grid[colIdx] {
			if grid[colIdx][rowIdx] == 0 {
				landIdx, isEdage := checkLandStatus(colIdx, rowIdx, grid, landConvertMap)
				if landIdx != -1 {
					grid[colIdx][rowIdx] = landIdx
				} else {
					grid[colIdx][rowIdx] = currentIsland
					currentIsland += 1
				}
				isNotIslandMap[grid[colIdx][rowIdx]] = isEdage || isNotIslandMap[grid[colIdx][rowIdx]]
			}
		}
	}

	for {
		flag := true
		for key := range landConvertMap {
			val := landConvertMap[key]

			vv, ok := landConvertMap[val]
			if ok {
				landConvertMap[key] = vv
				flag = false
			}
		}
		if flag {
			break
		}
	}

	for key := range landConvertMap {
		val := landConvertMap[key]
		isNotIslandMap[val] = isNotIslandMap[val] || isNotIslandMap[key]
		isNotIslandMap[key] = true
	}

	maxIdx := 1
	for colIdx := range grid {
		for rowIdx := range grid[colIdx] {
			val := grid[colIdx][rowIdx]
			target, ok := landConvertMap[val]
			if ok {
				val = target
			}
			if maxIdx < val {
				maxIdx = val
			}
			grid[colIdx][rowIdx] = val
		}
	}

	result := 0
	for idx := 2; idx <= maxIdx; idx++ {
		if isNotIslandMap[idx] == false {
			result += 1
		}
	}
	// fmt.Println(landConvertMap)
	// for _, i := range grid {
	// 	for _, j := range i {
	// 		fmt.Printf(" %2d", j)
	// 	}
	// 	fmt.Println()
	// }
	return result
}

func TestIsland(t *testing.T) {
	assert.Equal(t, 5,
		closedIsland([][]int{
			{0, 0, 1, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1, 0, 1},
			{0, 1, 1, 0, 0, 0, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 1, 1, 0},
			{0, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1},
			{1, 1, 1, 0, 0, 0, 1, 1, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0, 0},
			{0, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1},
			{1, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0},
			{0, 1, 1, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 0, 1, 0},
			{0, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 0, 0, 1, 1},
			{0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1},
			{0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 1, 1, 1},
			{0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 0, 1, 0, 0},
			{1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1},
			{0, 1, 1, 0, 1, 1, 1, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0},
		}), nil)
	assert.Equal(t, 6,
		closedIsland([][]int{
			{0, 1, 0, 1, 0, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1},
			{0, 1, 1, 0, 0, 0, 0, 1, 1, 1, 0, 1, 0, 1, 1, 0, 0, 1, 0, 1},
			{1, 1, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0},
			{0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0},
			{1, 1, 0, 0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 1, 1, 1, 0, 0, 1, 1},
			{1, 1, 0, 0, 0, 0, 0, 1, 0, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 1},
			{1, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 1, 1, 1, 0, 1, 0},
			{1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			{0, 0, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 0, 1, 1, 0, 1, 0, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 1, 1},
			{1, 0, 0, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 0, 0, 0, 0, 1, 1, 0},
			{1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 0, 1, 0, 1, 1, 1},
			{0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0},
			{1, 0, 0, 0, 1, 1, 0, 1, 1, 1, 0, 0, 1, 0, 1, 1, 0, 1, 0, 1},
		}), nil)

	assert.Equal(t, 0,
		closedIsland([][]int{
			{0, 1, 1, 1, 0, 1},
			{1, 0, 1, 0, 0, 1},
			{1, 0, 1, 0, 1, 1},
			{1, 0, 0, 0, 1, 1},
			{0, 1, 1, 1, 0, 0},
		}), nil)

	assert.Equal(t, 1,
		closedIsland([][]int{
			{0, 1, 1, 1, 0},
			{1, 0, 1, 0, 1},
			{1, 0, 1, 0, 1},
			{1, 0, 0, 0, 1},
			{0, 1, 1, 1, 0},
		}), nil)

	assert.Equal(t, 2,
		closedIsland([][]int{
			{1, 1, 1, 1, 1, 1, 1},
			{1, 0, 0, 0, 0, 0, 1},
			{1, 0, 1, 1, 1, 0, 1},
			{1, 0, 1, 0, 1, 0, 1},
			{1, 0, 1, 1, 1, 0, 1},
			{1, 0, 0, 0, 0, 0, 1},
			{1, 1, 1, 1, 1, 1, 1},
		}), nil)

	assert.Equal(t, 2,
		closedIsland([][]int{
			{1, 1, 1, 1, 1, 1, 1, 0},
			{1, 0, 0, 0, 0, 1, 1, 0},
			{1, 0, 1, 0, 1, 1, 1, 0},
			{1, 0, 0, 0, 0, 1, 0, 1},
			{1, 1, 1, 1, 1, 1, 1, 0},
		}), nil)

	assert.Equal(t, 1,
		closedIsland([][]int{
			{0, 0, 1, 0, 0},
			{0, 1, 0, 1, 0},
			{0, 1, 1, 1, 0},
		}), nil)

	assert.Equal(t, 0,
		closedIsland([][]int{
			{0, 0, 1, 0, 0},
			{0, 1, 0, 1, 0},
			{0, 1, 0, 1, 0},
			{0, 1, 0, 1, 0},
			{0, 0, 0, 1, 0},
			{0, 1, 1, 1, 0},
		}), nil)
}
