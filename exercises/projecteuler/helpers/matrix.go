package helpers

func Matrix(m int, n int, init int) [][]int {
	var matrix [][]int

	for i := 0; i < m; i++ {
		row := []int{}
		for j := 0; j < n; j++ {
			row = append(row, init)
		}
		matrix = append(matrix, row)
	}
	return matrix
}
