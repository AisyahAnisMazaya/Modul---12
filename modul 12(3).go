package main

import (
	"fmt"
)

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		// memindahkan elemen yang lebih besar dari key ke satu posisi di depan
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func hitungMedian(arr []int) int {
	n := len(arr)
	if n%2 == 1 { //jumlah data ganjil
		return arr[n/2]
	}
	// jumlah data genap, hitung rerata dua nilai tengah (dibulatkan ke bawah)
	return (arr[n/2-1] + arr[n/2]) / 2
}

func main() {
	const maksimalData = 1000000
	var data []int
	var angka int
	for {
		fmt.Scan(&angka)

		if angka == -5313 {
			break
		}

		if angka == 0 {
			InsertionSort(data)

			median := hitungMedian(data)
			fmt.Println(median)
		} else { // menambhakan bilangan ke array
			data = append(data, angka)
			if len(data) > maksimalData {
				fmt.Println("Data terlalu banyak, program dihentikan.")
				return
			}
		}
	}
}
