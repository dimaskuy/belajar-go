package main

import (
	"fmt"
	"math"
)

func main() {
	for {
		var input int
		fmt.Print("\n1. Program after-discount\n2. Program cari berat-badan di BMI\n3. Koordinat dari segitiga siku-siku terpanjang\n4. keluar \nMasukkan input: ")
		fmt.Scan(&input)
		// fmt.Print("hello")

		switch input {
		case 1:
			totalHarga()
		case 2:
			hitungBMI()
		case 3:
			koordinat()
		default:
			fmt.Print("error!")
		}
	}

}

func totalHarga() {
	var awal, diskon int
	fmt.Print("Masukkan harga awal dan diskon (contoh: 100000 10): ")
	fmt.Scan(&awal, &diskon)
	hasil := awal - (awal * diskon / 100)
	fmt.Print(hasil)
}

func hitungBMI() {
	var bmi, tb float64
	fmt.Print("Masukkan BMI (kg) dan tinggi (m) badan dalam desimal (contoh: 22.85 1.75): ")
	fmt.Scan(&bmi, &tb)

	bb := math.Ceil(bmi * (tb * tb))
	fmt.Print(bb)
}

func koordinat() {
	var ax, ay, bx, by, cx, cy float64
	fmt.Printf("Masukkan koordinat\nBaris pertama: koordinat A, Baris kedua: koordinat B, Baris ketiga: koordinat C\n")
	fmt.Scan(&ax, &ay)
	fmt.Scan(&bx, &by)
	fmt.Scan(&cx, &cy)

	ab := math.Sqrt(math.Pow(ax-bx, 2) + math.Pow(ay-by, 2))
	bc := math.Sqrt(math.Pow(cx-bx, 2) + math.Pow(cy-by, 2))
	ca := math.Sqrt(math.Pow(cx-ax, 2) + math.Pow(cy-ay, 2))

	max := math.Max(ab, math.Max(bc, ca))
	fmt.Printf("%.2f", max)

}
