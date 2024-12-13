package main

import "fmt"

const maxSize = 100

type datInt [maxSize]int

func pengurutanMembesar(A *datInt, N int) {
	for i := 0; i < N-1; i++ {
		minIdx := i
		for j := i + 1; j < N; j++ {
			if A[j] < A[minIdx] {
				minIdx = j
			}
		}
		A[i], A[minIdx] = A[minIdx], A[i]
	}
}

func main() {
	var n int
	fmt.Scan(&n) 

	if n <= 0 || n >= 1000000 || n > maxSize {
		fmt.Println("Jumlah daerah tidak memenuhi syarat")
		return
	}
	
	result := make([][]int, n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m) 

		if m <= 0 || m >= 1000000 || m > maxSize {
			fmt.Println("Jumlah rumah tidak memenuhi syarat")
			return
		}

		var houses datInt
		for j := 0; j < m; j++ {
			fmt.Scan(&houses[j])
		}

		pengurutanMembesar(&houses, m) 
		result[i] = houses[:m]
	}

	for i := 0; i < len(result); i++ {
		houses := result[i]
		for j := 0; j < len(houses); j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(houses[j])
		}
		fmt.Println()
	}
}