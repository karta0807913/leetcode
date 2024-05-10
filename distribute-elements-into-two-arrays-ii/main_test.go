package main

import (
	"fmt"
	"sort"

	rbt "github.com/emirpasic/gods/v2/trees/redblacktree"
)

func resultArray(nums []int) []int {
	var tree rbt.Tree
	var node rbt.Node
	greaterCount := func(arr []int, target int) int {
		sort.Ints(arr)
		return len(arr) - sort.Search(len(arr), func(i int) bool {
			return target < arr[i]
		})
	}
	arr1, arr2 := []int{nums[0]}, []int{nums[1]}
	_arr1, _arr2 := []int{nums[0]}, []int{nums[1]}
	for _, val := range nums[2:] {
		arr1Count := greaterCount(_arr1, val)
		arr2Count := greaterCount(_arr2, val)
		fmt.Println(arr1Count, arr2Count, val)
		if arr1Count > arr2Count {
			arr1 = append(arr1, val)
			_arr1 = append(_arr1, val)
		} else if arr1Count < arr2Count {
			arr2 = append(arr2, val)
			_arr2 = append(_arr2, val)
		} else {
			fmt.Println(arr1, arr2, val)
			if len(arr1) > len(arr2) {
				arr2 = append(arr2, val)
				_arr2 = append(_arr2, val)
			} else {
				arr1 = append(arr1, val)
				_arr1 = append(_arr1, val)
			}
		}
	}
	fmt.Println(arr1, arr2)
	return append(arr1, arr2...)
}
