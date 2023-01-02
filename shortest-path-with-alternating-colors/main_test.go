package main

type Color int

const (
	Red  Color = 0
	Blue Color = 1
)

type ColorEdge struct {
	NextNode int
	Color    Color
}

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	colorMap := make([][2]int, n)
	edgesMap := make(map[int][]ColorEdge)
	for _, edge := range redEdges {
		edgesMap[edge[0]] = append(edgesMap[edge[0]], ColorEdge{
			NextNode: edge[1],
			Color:    Red,
		})
	}
	for _, edge := range blueEdges {
		edgesMap[edge[0]] = append(edgesMap[edge[0]], ColorEdge{
			NextNode: edge[1],
			Color:    Blue,
		})
	}
	current := []int{0}
	next := make([]int, 0)
	cost := 1
	colorMap[0][0] = -1
	colorMap[0][1] = -1
	for len(current) != 0 {
		for _, id := range current {
			for _, edge := range edgesMap[id] {
				if edge.Color == Blue && colorMap[id][1] != 0 && colorMap[id][1] != cost && colorMap[edge.NextNode][0] == 0 {
					colorMap[edge.NextNode][0] = cost
					next = append(next, edge.NextNode)
				}
				if edge.Color == Red && colorMap[id][0] != 0 && colorMap[id][0] != cost && colorMap[edge.NextNode][1] == 0 {
					colorMap[edge.NextNode][1] = cost
					next = append(next, edge.NextNode)
				}
			}
		}
		cost++
		current, next = next, current
		next = next[:0]
	}
	result := make([]int, n)
	for i, n := range colorMap {
		result[i] = n[0]
		if result[i] == 0 || (n[1] != 0 && n[1] < result[i]) {
			result[i] = n[1]
		}
		if result[i] == 0 {
			result[i] = -1
		}
	}
	result[0] = 0
	return result
}
