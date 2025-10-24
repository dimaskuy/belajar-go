package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan 3 nilai: ")
	fmt.Scan(&n)

	n1 := n / 100
	n2 := (n / 10) % 10
	n3 := n % 10

	fmt.Print(((n1 * n1 * n1) + (n2 * n2 * n2) + (n3 * n3 * n3)) == n)
}

func main2() {
	var n, a, b int
	fmt.Printf("==PILIH MENU==\nMasukkan angka:\n1. Penjumlahan\n2.Pengurangan\nPilih: ")
	fmt.Scan(&n)
	fmt.Printf("\nMasukkan angka 1: ")
	fmt.Scan(&a)
	fmt.Printf("\nMasukkan angka 2: ")
	fmt.Scan(&b)

	if n == 1 {
		// penjumlahan(a, b)
		fmt.Print(a + b)
	} else if n == 2 {
		// pengurangan(a, b)
		fmt.Print(a - b)
	}
}

// func penjumlahan(a int, b int) {
// 	fmt.Print(a + b)
// }
// func pengurangan(a int, b int) {
// 	fmt.Print(a - b)
// }
