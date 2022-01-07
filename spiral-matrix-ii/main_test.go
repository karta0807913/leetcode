package main

func generateMatrix(n int) [][]int {
	c1, c2 := n, n-1
	matrix := make([][]int, n)
	for idx := range matrix {
		matrix[idx] = make([]int, n)
	}
	x, y, d1, d2 := -1, 0, 1, 1
	p1, p2 := &x, &y
	idx := 1
	for c1 != 0 {
		// fmt.Printf("c1: %v, c2: %v\n", c1, c2)
		for i := 0; i != c1; i++ {
			*p1 += d1
			// fmt.Printf("col: %v, row: %v, val: %v\n", y, x, matrix[y][x])
			matrix[y][x] = idx
			idx += 1
		}
		c1, c2, p1, p2, d1, d2 = c2, c1-1, p2, p1, d2, -d1
		// c1, c2 = c2, c1-1
		// p1, p2 = p2, p1
		// d1, d2 = d2, -d1
	}
	return matrix
}
