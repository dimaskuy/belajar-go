package main

import "fmt"

func main() {
	// deklarasi variabel "milisec" dalam integer
	var milisec int
	// input "milisec"
	fmt.Print("Masukkan waktu (ms): ")
	// simpan variabel-nya
	fmt.Scan(&milisec)
	// proses konversi dgn "milisec" di-casting ke float32 dahulu lalu dibagi dengan 1000.0, konversi di-assign ke variabel "sec"  
	sec := float32(milisec) / 1000.0
	// tampilkan hasil variabel "sec", dan bulatkan hasil
	fmt.Printf("Waktu dalam detik: %.1f", sec)
}