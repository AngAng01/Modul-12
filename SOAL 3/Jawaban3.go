package main

import "fmt"

const MaxSize = 100

type datInt [MaxSize]int

func calculateMedian(A *datInt, N int) float64 {
	if N%2 == 0 {
		return float64((*A)[N/2-1]+(*A)[N/2]) / 2.0
	}
	return float64((*A)[N/2])
}

func selectionSort(A *datInt, N int) {
	for i := 0; i < N-1; i++ {
		minIdx := i
		for j := i + 1; j < N; j++ {
			if (*A)[j] < (*A)[minIdx] {
				minIdx = j
			}
		}
		(*A)[i], (*A)[minIdx] = (*A)[minIdx], (*A)[i]
	}
}

func main() {
	var data datInt
	var size int
	var zeroCount int

	for {
		var input int
		fmt.Scan(&input)

		if input == -5313541 {
			break
		}

		if input == 0 {
			if size == 0 {
				fmt.Println("Error: Data kosong.")
			} else {
				median := calculateMedian(&data, size)
				fmt.Printf("Median saat ini: %.2f\n", median)
			}

			zeroCount++
			if zeroCount == 2 {
				break
			}

		} else {
			if size >= MaxSize {
				fmt.Println("Error: Array penuh.")
				continue
			}
			data[size] = input
			size++

			selectionSort(&data, size)
		}
	}
	fmt.Println("Selesai")
}
