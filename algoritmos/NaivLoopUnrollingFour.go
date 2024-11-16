package algoritmos

// NaivLoopUnrollingFour realiza la multiplicación de matrices de forma clásica con "loop unrolling" de 4
func NaivLoopUnrollingFour(A, B [][]int) [][]int {
	n := len(A) // Se asume que A es una matriz cuadrada
	C := make([][]int, n)
	for i := range C {
		C[i] = make([]int, n)
	}

	// Loop unrolling de 4
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum := 0
			for k := 0; k < n-3; k += 4 {
				sum += A[i][k]*B[k][j] +
					A[i][k+1]*B[k+1][j] +
					A[i][k+2]*B[k+2][j] +
					A[i][k+3]*B[k+3][j]
			}
			// Si el tamaño de n no es múltiplo de 4, completar el cálculo para los elementos restantes
			for k := n - (n % 4); k < n; k++ {
				sum += A[i][k] * B[k][j]
			}
			C[i][j] = sum
		}
	}
	return C
}
