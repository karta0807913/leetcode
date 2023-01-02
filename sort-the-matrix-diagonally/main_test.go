package main

import "sort"

type IntPointerArray []*int

func (i IntPointerArray) Len() int {
	return len(i)
}

func (arr IntPointerArray) Less(i, j int) bool {
	return *arr[i] < *arr[j]
}

func (arr IntPointerArray) Swap(i, j int) {
	*arr[i], *arr[j] = *arr[j], *arr[i]
}

func diagonalSort(mat [][]int) [][]int {
	startingPoint := make([]int, 0, len(mat)+len(mat[0]))
	rowNum := len(mat[0])
	for i := 0; i < len(mat); i++ {
		startingPoint = append(startingPoint, i*rowNum)
	}
	for i := 1; i < rowNum; i++ {
		startingPoint = append(startingPoint, i)
	}
	cache := make(IntPointerArray, 0, rowNum)
	for _, point := range startingPoint {
		x, y := point%rowNum, point/rowNum
		for x < rowNum && y < len(mat) {
			cache = append(cache, &mat[y][x])
			x++
			y++
		}
		sort.Sort(cache)
		cache = cache[:0]
	}
	return mat
}

/*
[11,25,66, 1,69, 7],
[23,55,17,45,15,52],
[75,31,36,44,58, 8],
[22,27,33,25,68, 4],
[84,28,14,11, 5,50]

[ 5,17,4 ,1 ,22, 7],
[11,11,25,45, 8,28],
[14,23,25,44,58,15],
[52,27,31,36,50,66],
[84,69,75,33,55,68]

[ 5,17, 4, 1,52, 7],
[11,11,25,45, 8,69],
[14,23,25,44,58,15],
[22,27,31,36,50,66],
[84,28,75,33,55,68]
*/
