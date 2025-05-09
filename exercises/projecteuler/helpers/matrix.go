package helpers

func Matrix(m int, n int, init int) [][]int {
	var matrix [][]int

	for range m {
		row := []int{}
		for range n {
			row = append(row, init)
		}
		matrix = append(matrix, row)
	}
	return matrix
}
