package main

import (
	"Proyecto_Golang/algoritmos"
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	// Definir tamaños de las matrices
	sizes := []int{2, 4, 8, 16, 32, 64, 128, 256}

	// Registrar tiempos de ejecución
	results := make(map[string]map[int]float64) // Mapa de mapas para almacenar los tiempos por algoritmo y tamaño

	// Inicializar el mapa de resultados
	for _, algoritmo := range []string{
		"NaivOnArray", "NaivLoopUnrollingTwo", "NaivLoopUnrollingFour", "WinogradOriginal", "WinogradScaled",
		"StrassenNaiv", "III.3 Sequential Block", "IV.3 Sequential Block", "V.3 Sequential Block", "V.4 Parallel Block"} {
		results[algoritmo] = make(map[int]float64)
	}

	// Ejecutar y registrar tiempos de ejecución para cada algoritmo
	for _, size := range sizes {
		// Generar matrices de prueba
		A, B := generateTestMatrices(size)

		// Guardar las matrices en archivos si no existen
		saveMatrixToFile("resultados/matrices", "matrizA_"+strconv.Itoa(size)+".txt", A)
		saveMatrixToFile("resultados/matrices", "matrizB_"+strconv.Itoa(size)+".txt", B)

		// Ejecutar los algoritmos y medir el tiempo

		// NaivOnArray
		start := time.Now()
		algoritmos.NaivOnArray(A, B)
		results["NaivOnArray"][size] = time.Since(start).Seconds()

		// NaivLoopUnrollingTwo
		start = time.Now()
		algoritmos.NaivLoopUnrollingTwo(A, B)
		results["NaivLoopUnrollingTwo"][size] = time.Since(start).Seconds()

		// NaivLoopUnrollingFour
		start = time.Now()
		algoritmos.NaivLoopUnrollingFour(A, B)
		results["NaivLoopUnrollingFour"][size] = time.Since(start).Seconds()

		// WinogradOriginal
		start = time.Now()
		algoritmos.WinogradOriginal(A, B)
		results["WinogradOriginal"][size] = time.Since(start).Seconds()

		// WinogradScaled
		start = time.Now()
		algoritmos.WinogradScaled(A, B)
		results["WinogradScaled"][size] = time.Since(start).Seconds()

		// StrassenNaiv
		start = time.Now()
		algoritmos.StrassenNaiv(A, B)
		results["StrassenNaiv"][size] = time.Since(start).Seconds()

		// III.3 Sequential Block
		start = time.Now()
		algoritmos.SequentialBlock(A, B)
		results["III.3 Sequential Block"][size] = time.Since(start).Seconds()

		// IV.3 Sequential Block
		start = time.Now()
		algoritmos.SequentialBlockIV(A, B)
		results["IV.3 Sequential Block"][size] = time.Since(start).Seconds()

		// V.3 Sequential Block
		start = time.Now()
		algoritmos.SequentialBlockV(A, B)
		results["V.3 Sequential Block"][size] = time.Since(start).Seconds()

		// V.4 Parallel Block
		start = time.Now()
		algoritmos.ParallelBlockV(A, B)
		results["V.4 Parallel Block"][size] = time.Since(start).Seconds()
	}

	// Guardar los resultados en archivos separados por tamaño
	saveResultsBySize(sizes, results)
}

func nextPowerOfTwo(n int) int {
	if n <= 0 {
		return 1 // Manejar casos donde n es 0 o negativo
	}
	nextPow := int(math.Pow(2, math.Ceil(math.Log2(float64(n)))))
	return nextPow
}

func generateTestMatrices(n int) ([][]int, [][]int) {
	n = nextPowerOfTwo(n) // Asegura que el tamaño sea una potencia de 2
	A := make([][]int, n)
	B := make([][]int, n)
	for i := 0; i < n; i++ {
		A[i] = make([]int, n)
		B[i] = make([]int, n)
		for j := 0; j < n; j++ {
			// Generar números aleatorios con al menos 6 dígitos
			A[i][j] = rand.Intn(900000) + 100000 // Números aleatorios entre 100000 y 999999
			B[i][j] = rand.Intn(900000) + 100000
		}
	}
	return A, B
}

func saveMatrixToFile(folder, filename string, matrix [][]int) {
	// Verificar si el archivo ya existe
	_, err := os.Stat(folder + "/" + filename)
	if err == nil {
		// El archivo ya existe, no hacer nada
		return
	}

	// Crear el archivo para guardar la matriz
	file, err := os.Create(folder + "/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Escribir la matriz en el archivo
	for i := 0; i < len(matrix); i++ {
		// Escribir cada fila de la matriz con números separados por espacios
		for j := 0; j < len(matrix[i]); j++ {
			// Asegurar que cada número tenga al menos 6 dígitos
			file.WriteString(fmt.Sprintf("%06d ", matrix[i][j]))
		}
		// Escribir salto de línea después de cada fila
		file.WriteString("\n")
	}
}

func saveResultsBySize(sizes []int, results map[string]map[int]float64) {
	// Para cada tamaño de matriz
	for _, size := range sizes {
		// Crear el archivo para guardar los resultados de ese tamaño específico
		filename := fmt.Sprintf("resultados/tiempos/tiempos_%d.csv", size)
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// Crear el escritor CSV
		writer := csv.NewWriter(file)
		defer writer.Flush()

		// Escribir la cabecera
		writer.Write([]string{"Algoritmo", "Tiempo Promedio (segundos)"})

		// Escribir los resultados de los algoritmos para este tamaño
		for algo, times := range results {
			if time, exists := times[size]; exists {
				writer.Write([]string{algo, fmt.Sprintf("%f", time)})
			}
		}
	}
}
