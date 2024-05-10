package main

import "sort"

type Point struct {
	X, Y int
}

type Query struct {
	Index int
	Val   int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func checkPoint(grid [][]int, point Point) bool {
	if point.Y < 0 || len(grid) <= point.Y {
		return false
	}
	if point.X < 0 || len(grid[point.Y]) <= point.X {
		return false
	}
	if grid[point.Y][point.X] < 0 {
		return false
	}
	return true
}

func maxPoints(grid [][]int, queries []int) []int {
	ans := make([]int, len(queries))
	structQueries := make([]Query, 0, len(queries))
	for i, query := range queries {
		structQueries = append(structQueries, Query{
			Index: i,
			Val:   query,
		})
	}
	sort.Slice(structQueries, func(i, j int) bool {
		return structQueries[i].Val < structQueries[j].Val
	})
	counter := 0

	current := []Point{{0, 0}}
	next := []Point{}
	nextQuery := []Point{}
	for _, query := range structQueries {
		for len(current) != 0 {
			for _, point := range current {
				val := &grid[point.Y][point.X]
				if *val < 0 {
					continue
				}
				if *val >= query.Val {
					nextQuery = append(nextQuery, point)
					continue
				}
				counter++
				*val = -(*val)
				for _, delta := range [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
					point := Point{
						Y: point.Y + delta[0],
						X: point.X + delta[1],
					}
					if checkPoint(grid, point) {
						next = append(next, point)
					}
				}
			}
			current, next = next, current
			next = next[:0]
		}
		ans[query.Index] = counter
		current, nextQuery = nextQuery, current
		nextQuery = nextQuery[:0]
	}
	return ans
}
