package main

import "fmt"

func main() {
	var warna string

	fmt.Print("Input warna favorit: ")
	fmt.Scan(&warna)

	switch warna {
	case "merah":
		fmt.Println("Warna keberanian")
	case "biru":
		fmt.Println("Warna ketenangan")
	case "hijau":
		fmt.Println("Warna kesuburan")
	case "kuning":
		fmt.Println("Warna keceriaan")
	default:
		fmt.Println("Warna tidak dikenal")
	}

	fmt.Println("End")
}
