package main

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	edgeMap := make([]bool, n)
	for _, edge := range edges {
		edgeMap[edge[1]] = true
	}
	result := make([]int, 0)
	for idx, val := range edgeMap {
		if !val {
			result = append(result, idx)
		}
	}
	return result
}
