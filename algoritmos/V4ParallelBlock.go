package algoritmos

import (
	"sync"
)

// ParallelBlockV realiza la multiplicación de matrices utilizando bloques en paralelo
func ParallelBlockV(A, B [][]int) [][]int {
	n := len(A)     // Se asume que A es una matriz cuadrada
	blockSize := 64 // Tamaño del bloque, puede ser ajustado según la arquitectura
	C := make([][]int, n)

	// Inicialización de la matriz de resultados
	for i := range C {
		C[i] = make([]int, n)
	}

	// Crear un WaitGroup para esperar a que todos los goroutines terminen
	var wg sync.WaitGroup

	// Bucle sobre los bloques
	for i := 0; i < n; i += blockSize {
		for j := 0; j < n; j += blockSize {
			for k := 0; k < n; k += blockSize {
				// Para cada bloque, lanzar un goroutine para procesarlo en paralelo
				wg.Add(1)
				go func(i, j, k int) {
					defer wg.Done()
					// Bucle sobre los elementos dentro de cada bloque
					for iBlock := i; iBlock < minK(i+blockSize, n); iBlock++ {
						for jBlock := j; jBlock < minK(j+blockSize, n); jBlock++ {
							sum := 0
							// Bucle interno para procesar el bloque k
							for kBlock := k; kBlock < minK(k+blockSize, n); kBlock++ {
								sum += A[iBlock][kBlock] * B[kBlock][jBlock]
							}
							// Acumulación en la matriz C
							C[iBlock][jBlock] += sum
						}
					}
				}(i, j, k)
			}
		}
	}

	// Esperar a que todos los goroutines terminen
	wg.Wait()
	return C
}

// min es una función auxiliar para obtener el mínimo entre dos números
func minK(a, b int) int {
	if a < b {
		return a
	}
	return b
}
