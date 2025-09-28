// inisiasi library
package main

import "fmt"

func main() {
	var f int	// variabel fahrenheit (bilangan bulat)
	var c float32	// variabel hasil konversi ke celsius (real, 32bit)

	fmt.Print("Masukkan suhu Farenheit: ")	// input user suhu fahrenheit
	fmt.Scan(&f)	// baca input dari user dan simpan di variabel "f"
	c = (5.0/9.0) * float32(f-32)	// konversi fahrenheit ke celsius berbentuk bilangan real
	// cetak hasilnya dan tampilkan dalam dua angka belakang koma
	fmt.Printf("Suhu dalam Celsius (C) = %.2f", c)
}