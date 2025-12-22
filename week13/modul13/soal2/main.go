package main

import "fmt"

func main() {
	var jumlahTanaman int
	fmt.Print("Masukkan Jumlah Tanaman: ")
	fmt.Scan(&jumlahTanaman)

	const needed = 4.0
	const kapasitasEmber = 3.0
	const bocor = 0.5

	isiTotal := 0
	siramTotal := 0

	airEmber := 0.0
	tanaman := 1
	airTanaman := 0.0

	for tanaman <= jumlahTanaman {
		if airEmber <= 0 { // isi jika kosong
			isiTotal++
			airEmber = kapasitasEmber
		}

		airEmber -= bocor // bocor setiap ke tanaman lain
		siramTotal++

		if airEmber < 0 {
			airEmber = 0
		}

		sisaNeeded := needed - airTanaman // siram tanaman tiap kebutuhan air

		if airEmber >= sisaNeeded {
			airEmber -= sisaNeeded
			airTanaman = 0
			tanaman++
		} else {
			airTanaman += airEmber
			airEmber = 0
		}
	}

	fmt.Println("Total pengisian ember:", isiTotal, "kali")
	fmt.Println("Total penyiraman:", siramTotal, "kali")
}
