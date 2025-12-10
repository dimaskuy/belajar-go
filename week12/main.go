package main

import "fmt"

func main() {
	// a := 1

	// for a < 10 {
	// 	fmt.Println("Day ke-", a)
	// 	a++
	// }
	// fmt.Println("Loop selesai!")

	var a int
	var b int

	fmt.Scan(&a, &b)

	for (a+b)%2 == 0 {
		fmt.Println("Hasil penjumlahan", a+b)
		fmt.Scan(&a, &b)
	}

	fmt.Println("Program selesai")
}
