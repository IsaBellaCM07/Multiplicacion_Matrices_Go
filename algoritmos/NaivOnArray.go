package algoritmos

func NaivOnArray(A, B [][]int) [][]int {
	n := len(A)
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				result[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return result
}
