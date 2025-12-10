package main

import "fmt"

func main() {
	var kode, durasi, tarifPerJam int // variabel input, tarif
	var tipe string                   // jenis kendaraan
	isValid := true                   // status validasi
	isValidDescA := ""                // pesan error durasi
	isValidDescB := ""                // pesan error kode kendaraan

	fmt.Print("Masukkan kode kendaraan (1-4): ") // input & baca kode
	fmt.Scan(&kode)

	fmt.Print("Masukkan durasi parkir (jam): ") // input & baca durasi
	fmt.Scan(&durasi)

	if durasi <= 0 { // validasi durasi
		isValid = false                                      // tandai tidak valid
		isValidDescA = "Durasi tidak boleh kurang dari nol!" // pesan error
	}

	switch kode {
	case 1:
		tipe = "Motor"     // kendaraan motor
		tarifPerJam = 2000 // tarif motor
	case 2:
		tipe = "Mobil"
		tarifPerJam = 5000
	case 3:
		tipe = "Truk"
		tarifPerJam = 8000
	case 4:
		tipe = "Bus"
		tarifPerJam = 10000
	default:
		isValid = false                                    // kode tidak valid
		isValidDescB = "Input kode kendaraan tidak valid!" // pesan error
	}

	if isValid { // jika input valid
		total := 0 // total pembayaran

		if durasi > 5 { // lebih dari 5 jam
			total += 5 * tarifPerJam                                 // 5 jam pertama tarif normal
			sisa := durasi - 5                                       // jam ekstra
			total += int(float64(sisa) * float64(tarifPerJam) * 1.5) // tarif 1.5x
		} else {
			total = durasi * tarifPerJam // <=5 jam tarif normal
		}

		fmt.Println("tipe Kendaraan:", tipe)
		fmt.Println("Durasi Parkir:", durasi, "jam")
		fmt.Println("Tarif per Jam: Rp", tarifPerJam)
		fmt.Println("Total Bayar: Rp", total)
	} else {
		fmt.Printf("ERROR:\n%s\n%s", isValidDescA, isValidDescB) // error-handling
	}
}
