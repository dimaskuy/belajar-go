package main

import "fmt"

func main() {
	var h string
	fmt.Print("Masukkan hari (senin, jumat, minggu): ")
	fmt.Scan(&h)
	// input senin, jumat minggu
	if h == "senin" || h == "Senin" {
		fmt.Println("Awal minggu semangat kerja!")
	} else if h == "jumat" || h == "Jumat" {
		fmt.Println("Menjelang Weekend!")
	} else if h == "minggu" || h == "Minggu" {
		fmt.Println("Hari libur istirahat!")
	} else {
		fmt.Println("Hari tak dikenali!")
	}
}
