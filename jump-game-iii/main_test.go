package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func canReach(arr []int, start int) bool {
	searched := make(map[int]bool)
	workingIdx := []int{start}
	appendSearched := func(point int) bool {
		if point < 0 || len(arr) <= point {
			return false
		}
		if arr[point] == 0 {
			return true
		}
		if !searched[point] {
			searched[point] = true
			workingIdx = append(workingIdx, point)
		}
		return false
	}

	for len(workingIdx) != 0 {
		idx := len(workingIdx)
		for _, val := range workingIdx[:idx] {
			if appendSearched(val - arr[val]) {
				return true
			}
			if appendSearched(val + arr[val]) {
				return true
			}
		}
		workingIdx = workingIdx[idx:]
	}
	return false
}

func TestReach(t *testing.T) {
	assert.True(t, canReach([]int{4, 2, 3, 0, 3, 1, 2}, 5), nil)
	assert.True(t, canReach([]int{4, 2, 3, 0, 3, 1, 2}, 0), nil)
	assert.False(t, canReach([]int{3, 0, 2, 1, 2}, 2), nil)
}

// // slow method
// const CAN_REACH = -1
// const CANNOT_REACH = 1
// const UNKNOW = 0

// func canReach(arr []int, start int) bool {
// 	answer := make([]int, len(arr))
// 	searchFlag := false
// 	for idx, val := range arr {
// 		if val == 0 {
// 			answer[idx] = CAN_REACH
// 			searchFlag = true
// 		}
// 	}
// 	if !searchFlag {
// 		return false
// 	}
// 	if answer[start] == CAN_REACH {
// 		return true
// 	}

// 	for range answer {
// 		changed := false
// 		for idx, steps := range arr {
// 			if answer[idx] != UNKNOW {
// 				continue
// 			}
// 			left := idx - steps
// 			right := idx + steps
// 			tmp := answer[idx]
// 			if left >= 0 {
// 				answer[idx] |= answer[left]
// 			}
// 			if right < len(answer) {
// 				answer[idx] |= answer[right]
// 			}
// 			if tmp != answer[idx] {
// 				changed = true
// 			}
// 		}
// 		// fmt.Println(answer)
// 		if !changed {
// 			break
// 		}
// 		if answer[start] != UNKNOW {
// 			break
// 		}
// 	}

// 	return answer[start] == CAN_REACH
// }

// func TestPointReach(t *testing.T) {
// 	assert.Equal(t, CAN_REACH, pointCanReach([]int{UNKNOW, CAN_REACH, CANNOT_REACH}, 0), nil)
// 	assert.Equal(t, CANNOT_REACH, pointCanReach([]int{CANNOT_REACH, CANNOT_REACH, UNKNOW}, 2), nil)
// 	assert.Equal(t, UNKNOW, pointCanReach([]int{CANNOT_REACH, UNKNOW, CAN_REACH}, 2), nil)
// }
