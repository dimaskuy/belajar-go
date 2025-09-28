package main

import "fmt"

func main() {
	var n1, n2, n3, total int	// variabel dengan value yang akan diinputkan & totalnya (bil bulat)
	var avg float64	// variabel hasil dari rata-rata (real, 64bit)
	fmt.Print("Masukkan tiga nilai: ")	// user input tiga nilai
	fmt.Scan(&n1, &n2, &n3)	// tiga nilai itu disimpan ke masing-masing variabel (n1-n3)
	total = n1 + n2 + n3	// menjumlahkan semua tiga nilai
	avg = float64(total) / 3.0	// kalkulasi rata-rata dengan total dibagi dengan 3

	fmt.Printf("Total = %d", total)	// cetak hasil total
	fmt.Printf("Rata-rata = %f", avg)	// cetak hasil rata-rata
}