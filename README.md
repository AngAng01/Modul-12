# MODUL DUABELAS
  ## SOAL 1
  Program di atas adalah program untuk mengurutkan nomor rumah di setiap daerah dalam urutan membesar (ascending). 
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt' dan 'math').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Konstanta maxSize, Konstanta ini mendefinisikan ukuran maksimum array (datInt), yang bernilai 100. Konstanta ini digunakan sebagai batas untuk jumlah rumah di setiap daerah.
      - Tipe Data datInt, Sebuah tipe data berupa array dengan ukuran tetap ([maxSize]int) yang digunakan untuk menyimpan daftar nomor rumah di setiap daerah.
      - Fungsi pengurutanMembesar, Fungsi ini mengimplementasikan algoritma selection sort untuk mengurutkan elemen-elemen array dalam urutan membesar (ascending order). Fungsi menerima referensi ke array A (dari tipe datInt) dan jumlah elemen N yang akan diurutkan.
      
   ## Code Explanation
   ```go
	const maxSize = 100
   ```
   Kode di atas adalah kode deklarasi konstanta yang menetapkan nilai tetap 100. Konstanta ini digunakan untuk menentukan batas maksimum ukuran array dalam program, memastikan bahwa jumlah elemen tidak melebihi kapasitas yang ditentukan. 
  
   ```go
	type datInt [maxSize]int
   ```
   Kode di atas adalah kode deklarasi tipe data baru.Tipe ini mendefinisikan datInt sebagai array dengan ukuran tetap sebesar maxSize (yang sebelumnya telah dideklarasikan sebagai konstanta). Elemen-elemen dalam array ini bertipe int.

   ```go
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
   ```
   Kode di atas adalah kode implementasi algoritma selection sort. Algoritma ini digunakan untuk mengurutkan elemen-elemen array dalam urutan membesar (ascending order). Fungsi pengurutanMembesar bekerja dengan cara memilih elemen terkecil dari bagian array yang belum diurutkan dan menukarnya dengan elemen pertama yang belum terurut, hingga seluruh array terurut.
   
   ```go
	var n int
	fmt.Scan(&n)
   ```
   Kode di atas adalah kode untuk membaca input dari pengguna. Kode ini mendeklarasikan variabel n bertipe int dan kemudian menggunakan fungsi fmt.Scan(&n) untuk mengambil nilai input dari pengguna melalui konsol dan menyimpannya dalam variabel 

   ```go
	if n <= 0 || n >= 1000000 || n > maxSize {
		fmt.Println("Jumlah daerah tidak memenuhi syarat")
		return
	}
   ```
   Kode di atas adalah kode validasi input, Kode ini memeriksa apakah nilai n yang dimasukkan memenuhi syarat atau tidak. Jika n lebih kecil dari atau sama dengan 0, lebih besar dari atau sama dengan 1.000.000, atau lebih besar dari nilai maksimum yang ditetapkan oleh konstanta maxSize, maka program akan mencetak pesan kesalahan "Jumlah daerah tidak memenuhi syarat" dan menghentikan eksekusi lebih lanjut dengan return. 

   ```go
	result := make([][]int, n)
   ```
   Kode di atas adalah kode untuk membuat slice dua dimensi. Dengan menggunakan fungsi make, kode ini mendeklarasikan sebuah slice bernama result yang memiliki panjang n dan menyimpan elemen-elemen yang merupakan slice bertipe int. 
   
   ```go
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
   ```
   Kode di atas adalah kode untuk memproses input dan mengurutkan data dalam array. Dalam kode ini, program pertama-tama menerima input jumlah daerah (n). Kemudian, untuk setiap daerah, program membaca jumlah rumah (m) dan memvalidasi apakah jumlah rumah tersebut memenuhi syarat. Jika jumlah rumah tidak valid, program akan mencetak pesan kesalahan dan menghentikan eksekusi. Setelah itu, program membaca nomor rumah untuk setiap daerah dan menyimpannya dalam array houses bertipe datInt. Array ini kemudian diurutkan menggunakan fungsi pengurutanMembesar, yang mengurutkan elemen dalam urutan membesar. Hasil yang telah diurutkan kemudian disalin ke dalam slice dua dimensi result.

   ```go
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
   ```
   Kode di atas adalah kode untuk menampilkan hasil yang telah diproses dan disimpan dalam slice dua dimensi result.


   ## SOAL 2
  Program di atas adalah program untuk mengurutkan nomor rumah berdasarkan kriteria ganjil dan genap, di mana urutan pengurutan berbeda untuk setiap jenis.
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt' dan 'math').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Deklarasi Konstanta MAX, Menyimpan nilai batas maksimum yang dapat diterima untuk jumlah daerah dan jumlah rumah, yaitu 1.000.000.
      - Fungsi pengurutanMembesar, Mengimplementasikan algoritma pengurutan selection sort untuk mengurutkan array dalam urutan menaik (dari kecil ke besar).
      - Fungsi pengurutanMengecil, Mengimplementasikan algoritma pengurutan selection sort untuk mengurutkan array dalam urutan menurun (dari besar ke kecil).
      
   ## Code Explanation
   ```go
	const MAX int = 1000000 
   ```
   Kode di atas adalah kode untuk mendeklarasikan sebuah konstanta bernama MAX dengan nilai 1.000.000 dan tipe data integer (int). 
  
   ```go
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
   ```
   Kode di atas adalah kode untuk mengurutkan array A secara menaik menggunakan algoritma selection sort. Fungsi ini menerima dua parameter: sebuah slice A yang berisi elemen-elemen yang akan diurutkan dan N, yaitu jumlah elemen dalam array.

   ```go
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
   ```
   Kode di atas adalah kode untuk mengurutkan array A secara menurun menggunakan algoritma selection sort. Fungsi ini menerima dua parameter: sebuah slice A yang berisi elemen-elemen yang akan diurutkan dan N, yaitu jumlah elemen dalam array.
   
   ```go
	var jumlahDaerah int
	fmt.Scan(&jumlahDaerah)
   ```
   Kode di atas adalah kode untuk membaca input integer yang diberikan oleh pengguna dan menyimpannya dalam variabel jumlahDaerah. 

   ```go
	if jumlahDaerah <= 0 || jumlahDaerah >= MAX {
		fmt.Println("Jumlah daerah tidak memenuhi syarat")
		return
	}
   ```
  Kode di atas adalah kode untuk memvalidasi input jumlah daerah. Bagian ini memastikan bahwa nilai yang dimasukkan oleh pengguna untuk jumlahDaerah memenuhi syarat tertentu. Jika nilai yang dimasukkan tidak sesuai dengan syarat, program akan menampilkan pesan kesalahan dan menghentikan eksekusi lebih lanjut.

   ```go
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
   ```
   Kode di atas adalah bagian dari program yang menangani proses pembacaan input dan pengolahan data terkait dengan jumlah rumah di setiap daerah. Pertama, program membaca jumlah daerah dan memeriksa apakah nilainya valid. Setelah itu, untuk setiap daerah, program membaca jumlah rumah dan nomor rumah yang ada. Kemudian, nomor rumah dibagi menjadi dua kategori: rumah dengan nomor ganjil dan rumah dengan nomor genap. Jika ada rumah genap, maka kedua kategori (ganjil dan genap) diurutkan secara menurun, menggunakan fungsi pengurutanMengecil. Jika tidak ada rumah genap, hanya rumah ganjil yang diurutkan secara membesar menggunakan fungsi pengurutanMembesar. Hasilnya, nomor rumah ganjil ditampilkan terlebih dahulu diikuti dengan nomor rumah genap. Program memastikan bahwa hasil output dipisahkan dengan spasi dan ditampilkan secara terformat untuk setiap daerah.


  ## SOAL 3
  Program di atas adalah program untuk menghitung median dari input yang dimasukkan oleh pengguna secara bertahap. 
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt' dan 'math').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Konstanta MaxSize, Digunakan untuk menentukan ukuran maksimum array yang digunakan untuk menyimpan data input.
      - Tipe data datInt, Tipe array dengan ukuran tetap, yang digunakan untuk menyimpan angka-angka yang dimasukkan oleh pengguna.
      - Fungsi calculateMedian, Fungsi untuk menghitung median dari array yang sudah diurutkan.
      - Fungsi insertSorted, Fungsi untuk menyisipkan angka baru ke dalam array dengan cara yang terurut, menggunakan teknik penyisipan langsung.
      
   ## Code Explanation
   ```go
	const MaxSize = 100
   ```
   Kode di atas adalah kode deklarasi konstanta yang menetapkan nilai tetap 100. Konstanta ini digunakan untuk menentukan batas maksimum ukuran array dalam program, memastikan bahwa jumlah elemen tidak melebihi kapasitas yang ditentukan. 
  
   ```go
	type datInt [MaxSize]int
   ```
   Kode di atas adalah kode deklarasi tipe data baru.Tipe ini mendefinisikan datInt sebagai array dengan ukuran tetap sebesar maxSize (yang sebelumnya telah dideklarasikan sebagai konstanta). Elemen-elemen dalam array ini bertipe int.

   ```go
	func calculateMedian(A *datInt, N int) float64 {
		if N%2 == 0 {
			return float64((*A)[N/2-1]+(*A)[N/2]) / 2.0
		}
		return float64((*A)[N/2])
	}
   ```
   Kode di atas adalah kode untuk menghitung median dari sebuah array yang sudah diurutkan. Fungsi calculateMedian menerima parameter berupa pointer ke array A dan panjang array N. 
   
   ```go
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
   ```
   Kode di atas adalah kode untuk memasukkan nilai baru ke dalam array yang sudah terurut, dengan menjaga urutan array tetap terjaga setelah nilai dimasukkan. Fungsi insertSorted menerima tiga parameter: pointer ke array A, pointer ke variabel N (yang menyimpan jumlah elemen dalam array), dan nilai value yang akan dimasukkan.

   ```go
	var data datInt
	var size int
	var zeroCount int 
   ```
   Kode di atas adalah kode untuk mendeklarasikan tiga variabel yang digunakan dalam program. Variabel data bertipe datInt, yang merupakan array untuk menyimpan data input dengan ukuran maksimum yang telah ditentukan. Variabel size bertipe int berfungsi untuk melacak jumlah elemen yang telah dimasukkan ke dalam array data. Sedangkan variabel zeroCount, juga bertipe int, digunakan untuk menghitung berapa kali input dengan nilai 0 dimasukkan, yang nantinya digunakan untuk menghentikan program setelah dua kali input 0 diterima.

   ```go
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
   ```
   Kode di atas adalah kode untuk mengelola input data dalam sebuah loop yang berjalan terus-menerus hingga kondisi tertentu tercapai. Program akan meminta input bilangan integer dari pengguna. Jika input yang diberikan adalah -5313541, loop akan dihentikan dengan break. Jika input adalah 0, program akan memeriksa apakah data yang dimasukkan kosong (yaitu, jika size == 0). Jika kosong, program akan mencetak pesan kesalahan "Error: Data kosong." Jika ada data, maka program akan menghitung dan menampilkan median dari data yang telah dimasukkan dengan menggunakan fungsi calculateMedian. Setiap kali input 0 dimasukkan, variabel zeroCount akan bertambah 1. Jika input 0 dimasukkan dua kali, program akan keluar dari loop dan mencetak "Selesai". Untuk input selain 0, data akan disisipkan dan diurutkan menggunakan fungsi insertSorted.
   
   ```go
	fmt.Println("Selesai")
   ```
   Kode di atas adalah kode perintah untuk mencetak teks "Selesai" ke layar. Ini digunakan untuk menandakan bahwa program telah selesai menjalankan semua operasi dan keluar dari loop atau proses utama. 

   
