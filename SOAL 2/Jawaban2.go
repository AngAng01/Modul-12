package main

import "fmt"

const MAX int = 1000000 

func pengurutanMembesar(A []int, N int) {
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

func pengurutanMengecil(A []int, N int) {
	for i := 0; i < N-1; i++ {
		maxIdx := i
		for j := i + 1; j < N; j++ {
			if A[j] > A[maxIdx] {
				maxIdx = j
			}
		}
		A[i], A[maxIdx] = A[maxIdx], A[i]
	}
}

func main() {
	var jumlahDaerah int
	fmt.Scan(&jumlahDaerah)

	if jumlahDaerah <= 0 || jumlahDaerah >= MAX {
		fmt.Println("Jumlah daerah tidak memenuhi syarat")
		return
	}

	for i := 0; i < jumlahDaerah; i++ {
		var jumlahRumah int
		fmt.Scan(&jumlahRumah)

		if jumlahRumah <= 0 || jumlahRumah >= MAX {
			fmt.Println("Jumlah rumah tidak memenuhi syarat")
			return
		}

		nomorRumah := make([]int, jumlahRumah)
		for j := 0; j < jumlahRumah; j++ {
			fmt.Scan(&nomorRumah[j])
		}

		var ganjil, genap []int
		for j := 0; j < jumlahRumah; j++ {
			if nomorRumah[j]%2 == 0 {
				genap = append(genap, nomorRumah[j])
			} else {
				ganjil = append(ganjil, nomorRumah[j])
			}
		}

		if len(genap) > 0 {
			pengurutanMengecil(ganjil, len(ganjil)) 
			pengurutanMengecil(genap, len(genap))  
		} else {
			pengurutanMembesar(ganjil, len(ganjil)) 
		}

		fmt.Print("Hasil : ")
		for j := 0; j < len(ganjil); j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(ganjil[j])
		}
		for j := 0; j < len(genap); j++ {
			if len(ganjil) > 0 || j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(genap[j])
		}
		fmt.Println()
	}
}
