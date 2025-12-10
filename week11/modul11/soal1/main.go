package main

import "fmt"

func main() {
	var n int

	fmt.Print("Masukkan nilai akhir: ")
	fmt.Scan(&n)

	var result string

	switch n / 10 {
	case 9, 10:
		result = "A"
	case 8:
		result = "A"
	case 7:
		if n >= 75 {
			result = "B"
		} else {
			result = "C"
		}
	case 6:
		if n >= 65 {
			result = "C"
		} else {
			result = "D"
		}
	case 5:
		result = "D"
	default:
		result = "E"
	}

	fmt.Printf("Grade Anda: %s\n", result)
}
