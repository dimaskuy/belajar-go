package main

import "fmt"

func main() {
	var tipe string
	var durasi, total int
	fmt.Print("Kendaraan: ")
	fmt.Scan(&tipe)
	fmt.Print("Durasi: ")
	fmt.Scan(&durasi)

	switch tipe {
	case "mobil", "MOBIL":
		if durasi <= 2 && durasi >= 0 {
			total = 5000 * durasi
		} else if durasi > 2 {
			total = 4000 * durasi
		} else {
			fmt.Print("ERROR JIR!")
			break
		}
	case "motor", "MOTOR":
		if durasi <= 3 && durasi >= 0 {
			total = 2000 * durasi
		} else if durasi > 3 {
			total = 1500 * durasi
		} else {
			fmt.Print("ERROR JIR!")
			break
		}
	case "truk", "TRUK":
		if durasi <= 1 && durasi >= 0 {
			total = 10000 * durasi
		} else if durasi > 1 {
			total = 8000 * durasi
		} else {
			fmt.Print("ERROR JIR!")
			break
		}
	}

	if durasi > 8 {
		totalDiskon := float32(total) - (float32(total) * 0.1)
		fmt.Print("Total: ", totalDiskon)
	} else {
		fmt.Print("Total: ", total)
	}
}
