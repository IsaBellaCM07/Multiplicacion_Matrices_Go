package algoritmos

func WinogradOriginal(A, B [][]int) [][]int {
	n := len(A)
	C := make([][]int, n)
	for i := range C {
		C[i] = make([]int, n)
	}

	rowFactors := make([]int, n)
	colFactors := make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			rowFactors[i] += A[i][2*j] * A[i][2*j+1]
		}
	}

	for j := 0; j < n; j++ {
		for i := 0; i < n/2; i++ {
			colFactors[j] += B[2*i][j] * B[2*i+1][j]
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			C[i][j] = -rowFactors[i] - colFactors[j]
			for k := 0; k < n/2; k++ {
				C[i][j] += (A[i][2*k] + B[2*k+1][j]) * (A[i][2*k+1] + B[2*k][j])
			}
		}
	}

	if n%2 == 1 {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				C[i][j] += A[i][n-1] * B[n-1][j]
			}
		}
	}

	return C
}
