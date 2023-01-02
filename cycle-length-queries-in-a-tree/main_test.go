package main

func cycleLengthQueries(n int, queries [][]int) []int {
	answer := make([]int, 0, len(queries))
	for _, query := range queries {
		leftLayer, rightLayer := 0, 0
		left, right := query[0], query[1]
		if left < right {
			right, left = left, right
		}
		for left != right {
			left /= 2
			leftLayer++
			if left < right {
				leftLayer, rightLayer = rightLayer, leftLayer
				left, right = right, left
			}
		}
		answer = append(answer, leftLayer+rightLayer+2)
	}
	return answer
}
