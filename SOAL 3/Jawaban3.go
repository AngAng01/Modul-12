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

func insertSorted(A *datInt, N *int, value int) {
	if *N >= MaxSize {
		fmt.Println("Error: Array penuh.")
		return
	}
	(*A)[*N] = value
	*N++

	for i := *N - 1; i > 0 && (*A)[i] < (*A)[i-1]; i-- {
		(*A)[i], (*A)[i-1] = (*A)[i-1], (*A)[i]
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
			insertSorted(&data, &size, input)
		}
	}
	fmt.Println("Selesai")
}
