package main

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

type HeapSlice []*State

func (slice HeapSlice) Len() int            { return len(slice) }
func (slice HeapSlice) Less(i, j int) bool  { return slice[i].Sum < slice[j].Sum }
func (slice *HeapSlice) Push(s interface{}) { *slice = append(*slice, s.(*State)) }
func (slice HeapSlice) Swap(i, j int)       { slice[i], slice[j] = slice[j], slice[i] }
func (slice *HeapSlice) Pop() interface{} {
	n := len(*slice) - 1
	t := (*slice)[n]
	*slice = (*slice)[:n]
	return t
}

type State struct {
	leftIdx, rightIdx int
	Sum               int
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) (ans [][]int) {
	h := make(HeapSlice, 0)
	var nextState *State
	heap.Push(&h, &State{
		leftIdx:  0,
		rightIdx: 0,
		Sum:      nums1[0] + nums2[0],
	})
	if len(nums1) != 1 {
		nextState = &State{
			leftIdx:  1,
			rightIdx: 0,
			Sum:      nums1[1] + nums2[0],
		}
	}
	for ; k > 0; k-- {
		ans = append(ans, []int{nums1[h[0].leftIdx], nums2[h[0].rightIdx]})
		h[0].rightIdx += 1
		if h[0].rightIdx == len(nums2) {
			heap.Pop(&h)
		} else {
			h[0].Sum = nums1[h[0].leftIdx] + nums2[h[0].rightIdx]
			heap.Fix(&h, 0)
		}
		if len(h) == 0 && nextState == nil {
			return
		} else if nextState != nil && (len(h) == 0 || h[0].Sum > nextState.Sum) {
			heap.Push(&h, nextState)
			if nextState.leftIdx+1 == len(nums1) {
				nextState = nil
			} else {
				nextState = &State{
					leftIdx:  nextState.leftIdx + 1,
					rightIdx: 0,
					Sum:      nums1[nextState.leftIdx+1] + nums2[0],
				}
			}
		}
	}
	return
}

func TestPairs(t *testing.T) {
	assert.Equal(t, [][]int{
		{1, 2}, {1, 4}, {1, 6},
	}, kSmallestPairs([]int{
		1, 7, 11,
	}, []int{
		2, 4, 6,
	}, 3))
}
