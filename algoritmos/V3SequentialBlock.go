package algoritmos

// SequentialBlockV realiza la multiplicación de matrices utilizando bloques secuenciales
func SequentialBlockV(A, B [][]int) [][]int {
	n := len(A)     // Se asume que A es una matriz cuadrada
	blockSize := 64 // Tamaño del bloque, puede ser ajustado según la arquitectura
	C := make([][]int, n)

	// Inicialización de la matriz de resultados
	for i := range C {
		C[i] = make([]int, n)
	}

	// Bucle sobre los bloques
	for i := 0; i < n; i += blockSize {
		for j := 0; j < n; j += blockSize {
			for k := 0; k < n; k += blockSize {
				// Bucle sobre los elementos dentro de cada bloque
				for iBlock := i; iBlock < minJ(i+blockSize, n); iBlock++ {
					for jBlock := j; jBlock < minJ(j+blockSize, n); jBlock++ {
						sum := 0
						// Bucle interno para procesar el bloque k
						for kBlock := k; kBlock < minJ(k+blockSize, n); kBlock++ {
							sum += A[iBlock][kBlock] * B[kBlock][jBlock]
						}
						// Acumulación en la matriz C
						C[iBlock][jBlock] += sum
					}
				}
			}
		}
	}
	return C
}

// min es una función auxiliar para obtener el mínimo entre dos números
func minJ(a, b int) int {
	if a < b {
		return a
	}
	return b
}
