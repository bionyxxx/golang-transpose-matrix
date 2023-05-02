package main

import "fmt"

func main() {
	var rows, cols int

	// meminta ukuran matriks dari pengguna
	fmt.Print("Masukkan jumlah baris: ")
	_, err := fmt.Scan(&rows)
	if err != nil {
		panic(err)
	}
	fmt.Print("Masukkan jumlah kolom: ")
	_, err = fmt.Scan(&cols)
	if err != nil {
		panic(err)
	}

	// membuat matriks
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}

	// meminta inputan nilai pada matriks dari pengguna
	fmt.Println("Masukkan nilai pada matriks:")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("matrix[%d][%d] = ", i, j)
			_, err := fmt.Scan(&matrix[i][j])
			if err != nil {
				panic(err)
			}
		}
	}

	// membuat matriks transpose
	transpose := make([][]int, cols)
	for i := range transpose {
		transpose[i] = make([]int, rows)
	}

	// mengisi nilai pada matriks transpose
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transpose[j][i] = matrix[i][j]
		}
	}

	// mencetak matriks asli
	fmt.Println("Matriks Asli:")
	for i := 0; i < rows; i++ {
		fmt.Println(matrix[i])
	}

	// mencetak matriks transpose
	fmt.Println("Matriks Transpose:")
	for i := 0; i < cols; i++ {
		fmt.Println(transpose[i])
	}
}
