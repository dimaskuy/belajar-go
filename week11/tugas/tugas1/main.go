package main

import "fmt"

func main() {
	var harga float64
	var mata, kategori string

	fmt.Print("Masukkan harga: ")
	fmt.Scan(&harga)

	fmt.Print("Masukkan mata uang (USD/EUR/JPY/IDR): ")
	fmt.Scan(&mata)

	fmt.Print("Masukkan kategori pelanggan (REGULAR/MEMBER/VIP): ")
	fmt.Scan(&kategori)

	var diskon float64
	switch mata {
	case "USD", "usd":
		if harga > 1000 {
			diskon = 0.05
		}
	case "EUR", "eur":
		if harga > 900 {
			diskon = 0.07
		}
	case "JPY", "jpy":
		if harga > 120000 {
			diskon = 0.03
		}
	case "IDR", "idr":
		if harga > 10000000 {
			diskon = 0.02
		}
	default:
		fmt.Println("Mata uang tidak valid!")
		return
	}

	var pajak float64
	switch kategori {
	case "regular", "REGULAR":
		pajak = 0.10
	case "member", "MEMBER":
		pajak = 0.05
	case "VIP", "vip":
		pajak = 0.00
	default:
		fmt.Println("Kategori pelanggan tidak valid!")
		return
	}
	hargaSetelahDiskon := harga - (harga * diskon)
	total := hargaSetelahDiskon + pajak

	fmt.Printf("Total: %.2f %s\n", total, mata)
}
