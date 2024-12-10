package main

import (
	"fmt"
)

func SelectionSort(arr *[]int) {
	for i := 0; i < len(*arr)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(*arr); j++ {
			if (*arr)[j] < (*arr)[minIdx] {
				minIdx = j
			}
		}
		(*arr)[i], (*arr)[minIdx] = (*arr)[minIdx], (*arr)[i]
	}
}

func main() {
	const maksimalN = 1000
	const maksimalM = 1000000
	var m int
	var n int
	fmt.Scan(&n)
	if n <= 0 || n > maksimalN {
		fmt.Println("Jumlah daerah harus antara 1 dan", maksimalN)
		return
	}

	// Deklarasi array untuk menyimpan nomor rumah dan jumlah rumah di setiap daerah
	var rmhKerabat = make([][]int, n)
	var panjangDaerah = make([]int, n)

	// Proses setiap daerah
	for i := 0; i < n; i++ {
		fmt.Scan(&m) // Membaca jumlah rumah kerabat di daerah tersebut
		if m <= 0 || m > maksimalM {
			fmt.Println("Jumlah rumah di daerah", i+1, "harus antara 1 dan", maksimalM)
			continue // Validasi batas m
		}
		panjangDaerah[i] = m // Simpan jumlah rumah di daerah i
		// Membaca nomor rumah kerabat dan simpan dalam array
		rmhKerabat[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&rmhKerabat[i][j])
		}
		SelectionSort(&rmhKerabat[i])
	}

	for i := 0; i < n; i++ {
		for j := 0; j < panjangDaerah[i]; j++ {
			fmt.Printf("%d ", rmhKerabat[i][j])
		}
		fmt.Println()
	}
}
