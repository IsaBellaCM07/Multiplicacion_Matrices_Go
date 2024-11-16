package algoritmos

// WinogradScaled realiza la multiplicación de matrices usando el algoritmo Winograd Scaled
func WinogradScaled(A, B [][]int) [][]int {
	n := len(A) // Se asume que A es una matriz cuadrada
	C := make([][]int, n)
	for i := range C {
		C[i] = make([]int, n)
	}

	// Precomputación de las filas y columnas
	rowFactor := make([]int, n)
	colFactor := make([]int, n)

	for i := 0; i < n; i++ {
		rowFactor[i] = 0
		for k := 0; k < n-1; k += 2 {
			rowFactor[i] += A[i][k] * A[i][k+1]
		}
	}

	for j := 0; j < n; j++ {
		colFactor[j] = 0
		for k := 0; k < n-1; k += 2 {
			colFactor[j] += B[k][j] * B[k+1][j]
		}
	}

	// Cálculo de la multiplicación con la optimización de Winograd
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum := -rowFactor[i] - colFactor[j]
			for k := 0; k < n-1; k += 2 {
				sum += (A[i][k] + B[k+1][j]) * (A[i][k+1] + B[k][j])
			}
			// Si n es impar, añadir el último término
			if n%2 != 0 {
				sum += A[i][n-1] * B[n-1][j]
			}
			C[i][j] = sum
		}
	}
	return C
}
