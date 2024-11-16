package algoritmos

// NaivLoopUnrollingTwo realiza la multiplicación de matrices de forma clásica con "loop unrolling" de 2
func NaivLoopUnrollingTwo(A, B [][]int) [][]int {
	n := len(A) // Se asume que A es una matriz cuadrada
	C := make([][]int, n)
	for i := range C {
		C[i] = make([]int, n)
	}

	// Loop unrolling de 2
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum := 0
			for k := 0; k < n-1; k += 2 {
				sum += A[i][k]*B[k][j] + A[i][k+1]*B[k+1][j]
			}
			// Si el tamaño de n es impar, completar el cálculo para el último elemento
			if n%2 != 0 {
				sum += A[i][n-1] * B[n-1][j]
			}
			C[i][j] = sum
		}
	}
	return C
}
