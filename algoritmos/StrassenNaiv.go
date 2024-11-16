package algoritmos

import (
	"math"
)

// Algoritmo de Strassen (Naiv) corregido
func StrassenNaiv(A, B [][]int) [][]int {
	n := len(A)

	// Rellenar matrices a la siguiente potencia de 2
	originalSize := n
	A = padMatrix(A)
	B = padMatrix(B)
	n = len(A)

	if n == 1 {
		return [][]int{{A[0][0] * B[0][0]}}
	}

	mid := n / 2

	A11 := subMatrix(A, 0, 0, mid)
	A12 := subMatrix(A, 0, mid, mid)
	A21 := subMatrix(A, mid, 0, mid)
	A22 := subMatrix(A, mid, mid, mid)

	B11 := subMatrix(B, 0, 0, mid)
	B12 := subMatrix(B, 0, mid, mid)
	B21 := subMatrix(B, mid, 0, mid)
	B22 := subMatrix(B, mid, mid, mid)

	M1 := StrassenNaiv(addMatrix(A11, A22), addMatrix(B11, B22))
	M2 := StrassenNaiv(addMatrix(A21, A22), B11)
	M3 := StrassenNaiv(A11, subtractMatrix(B12, B22))
	M4 := StrassenNaiv(A22, subtractMatrix(B21, B11))
	M5 := StrassenNaiv(addMatrix(A11, A12), B22)
	M6 := StrassenNaiv(subtractMatrix(A21, A11), addMatrix(B11, B12)) // Corregido
	M7 := StrassenNaiv(subtractMatrix(A12, A22), addMatrix(B21, B22)) // Corregido

	C11 := addMatrix(subtractMatrix(addMatrix(M1, M4), M5), M7)
	C12 := addMatrix(M3, M5)
	C21 := addMatrix(M2, M4)
	C22 := addMatrix(subtractMatrix(addMatrix(M1, M3), M2), M6)

	C := make([][]int, n)
	for i := 0; i < n; i++ {
		C[i] = make([]int, n)
	}

	for i := 0; i < mid; i++ {
		for j := 0; j < mid; j++ {
			C[i][j] = C11[i][j]
			C[i][j+mid] = C12[i][j]
			C[i+mid][j] = C21[i][j]
			C[i+mid][j+mid] = C22[i][j]
		}
	}

	// Quitar el relleno de la matriz resultante
	C = unpadMatrix(C, originalSize)

	return C
}

// Función para rellenar una matriz a la siguiente potencia de 2
func padMatrix(matrix [][]int) [][]int {
	n := len(matrix)
	nextPowerOfTwo := int(math.Pow(2, math.Ceil(math.Log2(float64(n)))))

	if n == nextPowerOfTwo {
		return matrix
	}

	paddedMatrix := make([][]int, nextPowerOfTwo)
	for i := 0; i < nextPowerOfTwo; i++ {
		paddedMatrix[i] = make([]int, nextPowerOfTwo)
		for j := 0; j < nextPowerOfTwo; j++ {
			if i < n && j < n {
				paddedMatrix[i][j] = matrix[i][j]
			} else {
				paddedMatrix[i][j] = 0
			}
		}
	}
	return paddedMatrix
}

// Función para quitar el relleno de una matriz
func unpadMatrix(matrix [][]int, originalSize int) [][]int {
	unpaddedMatrix := make([][]int, originalSize)
	for i := 0; i < originalSize; i++ {
		unpaddedMatrix[i] = make([]int, originalSize)
		copy(unpaddedMatrix[i], matrix[i][:originalSize])
	}
	return unpaddedMatrix
}

// Función para sumar dos matrices
func addMatrix(A, B [][]int) [][]int {
	n := len(A)
	C := make([][]int, n)
	for i := range C {
		C[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			C[i][j] = A[i][j] + B[i][j]
		}
	}
	return C
}

// Función para restar dos matrices
func subtractMatrix(A, B [][]int) [][]int {
	n := len(A)
	C := make([][]int, n)
	for i := range C {
		C[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			C[i][j] = A[i][j] - B[i][j]
		}
	}
	return C
}

// Función para obtener una submatriz
func subMatrix(matrix [][]int, startRow, startCol, size int) [][]int {
	subMatrix := make([][]int, size)
	for i := 0; i < size; i++ {
		subMatrix[i] = make([]int, size)
		for j := 0; j < size; j++ {
			subMatrix[i][j] = matrix[startRow+i][startCol+j]
		}
	}
	return subMatrix
}
