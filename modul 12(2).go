package main

import (
	"fmt"
)

func slcSortAsc(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func slcSortDsc(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		maxIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
	}
}

func main() {
	const maxN = 1000
	const maxM = 1000000
	var m int
	var n int
	fmt.Scan(&n)
	if n <= 0 || n >= maxN {
		fmt.Println("Jumlah daerah harus antara 1 dan", maxN)
		return
	}

	rmhKerabat := make([][]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&m) // menginputkan jumlah rumah kerabat di daerah tersebut
		if m <= 0 || m >= maxM {
			fmt.Println("Jumlah rumah di daerah", i+1, "harus antara 1 dan", maxM)
			return
		}

		// Array sementara untuk menyimpan nomor rumah ganjil dan genap
		var ganjil, genap []int

		for j := 0; j < m; j++ {
			var nomor int
			fmt.Scan(&nomor)
			if nomor%2 == 1 { // Ganjil
				ganjil = append(ganjil, nomor)
			} else { // Genap
				genap = append(genap, nomor)
			}
		}

		//ganjil menaik
		slcSortAsc(ganjil)

		//genap menurun
		slcSortDsc(genap)

		//ganjil dan genap, append digunakan untuk menggabungkan array ganjil genap
		rmhKerabat[i] = append(ganjil, genap...)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < len(rmhKerabat[i]); j++ {
			fmt.Printf("%d ", rmhKerabat[i][j])
		}
		fmt.Println()
	}
}
