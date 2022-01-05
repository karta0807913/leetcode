package main

func rotate(matrix [][]int) {
	for boundary := 0; len(matrix)-1 > boundary*2; boundary++ {
		for idx := boundary; idx < len(matrix[boundary])-boundary-1; idx++ {
			// fmt.Printf("(%d, %d), (%d, %d), (%d, %d)\n", boundary, idx, idx, len(matrix)-boundary-1, len(matrix)-boundary-1, len(matrix)-idx-1)
			matrix[idx][len(matrix)-boundary-1],
				matrix[len(matrix)-boundary-1][len(matrix)-idx-1],
				matrix[len(matrix)-idx-1][boundary],
				matrix[boundary][idx] =
				matrix[boundary][idx],
				matrix[idx][len(matrix)-boundary-1],
				matrix[len(matrix)-boundary-1][len(matrix)-idx-1],
				matrix[len(matrix)-idx-1][boundary]
		}
	}
}
