package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if b != 0 {
		hitung := 0

		for a >= b {
			a = a - b
			hitung++
		}

		fmt.Println(hitung)
	} else {
		fmt.Println("ERROR: Tidak bisa bagi dengan nol")
	}
}
