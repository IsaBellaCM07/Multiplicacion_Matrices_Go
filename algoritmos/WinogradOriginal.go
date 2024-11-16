package algoritmos

func WinogradOriginal(A, B [][]int) [][]int {
	n := len(A)
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	rowFactor := make([]int, n)
	colFactor := make([]int, n)

	for i := 0; i < n; i++ {
		rowFactor[i] = A[i][0] * B[0][i]
		colFactor[i] = A[i][0] * B[0][i]
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result[i][j] = rowFactor[i] + colFactor[j]
		}
	}
	return result
}
