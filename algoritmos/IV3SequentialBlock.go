package algoritmos

// SequentialBlockIV realiza la multiplicación de matrices utilizando bloques secuenciales
func SequentialBlockIV(A, B [][]int) [][]int {
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
				for iBlock := i; iBlock < minI(i+blockSize, n); iBlock++ {
					for jBlock := j; jBlock < minI(j+blockSize, n); jBlock++ {
						sum := 0
						// Bucle interno para procesar el bloque k
						for kBlock := k; kBlock < minI(k+blockSize, n); kBlock++ {
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

// minI es una función auxiliar para obtener el mínimo entre dos números
func minI(a, b int) int {
	if a < b {
		return a
	}
	return b
}
