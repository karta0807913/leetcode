package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type IntStruct struct {
	data []int
	Get  func(x int) int
}

func (x IntStruct) Less(i, j int) bool {
	a := x.Get(x.data[i])
	b := x.Get(x.data[j])
	if a == b {
		return x.data[i] < x.data[j]
	}
	return a < b
}

func (x IntStruct) Len() int {
	return len(x.data)
}

func (p IntStruct) Search(x int) int {
	return sort.IntSlice(p.data).Search(x)
}

func (x IntStruct) Sort() {
	sort.Sort(x)
}

func (x IntStruct) Swap(i, j int) {
	sort.IntSlice(x.data).Swap(i, j)
}

func getKth(lo int, hi int, k int) int {
	result := IntStruct{
		data: make([]int, hi-lo+1),
	}
	for idx := 0; idx < hi-lo+1; idx++ {
		result.data[idx] = lo + idx
	}
	// storage := make(map[int]int)
	storage := [250505]uint8{}
	storage[1] = 0
	result.Get = func(x int) int {
		if x == 1 {
			return 1
		}
		val := int(storage[x])
		if val == 0 {
			t := 0
			if x&1 == 1 {
				t = x*3 + 1
			} else {
				t = x / 2
			}
			val = result.Get(t) + 1
			storage[x] = uint8(val)
		}
		return val
	}
	result.Sort()
	return result.data[k-1]
}

func normalTest(x int) int {
	count := 0
	for x != 1 {
		if x&1 == 1 {
			x = x*3 + 1
		} else {
			x /= 2
		}
		count += 1
	}
	return count
}

func TestGet(t *testing.T) {
	assert.Equal(t, 1, getKth(1, 1, 1), nil)
	assert.Equal(t, 13, getKth(12, 15, 2), nil)
	assert.Equal(t, 7, getKth(7, 11, 4), nil)
	assert.Equal(t, 13, getKth(10, 20, 5), nil)
	assert.Equal(t, 570, getKth(1, 1000, 777), nil)
	result := IntStruct{}
	storage := make(map[int]int)
	storage[1] = 0
	result.Get = func(x int) int {
		val, ok := storage[x]
		if !ok {
			t := 0
			if x&1 == 1 {
				t = x*3 + 1
			} else {
				t = x / 2
			}
			val = result.Get(t) + 1
			storage[x] = val
		}
		return val
	}
	for i := 1; i <= 1000; i++ {
		assert.Equal(t, normalTest(i), result.Get(i), nil)
	}
}
