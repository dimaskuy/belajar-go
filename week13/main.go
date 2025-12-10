package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var tebak int
	a := rand.Intn(10)
	fmt.Print(a)

	for tebak != a {
		fmt.Print("\nMasukkan tebakan (1-10): ")
		fmt.Scanln(&tebak)
		if tebak == a-1 || tebak == a+1 || tebak == a-2 || tebak == a+2 {
			fmt.Print("Salah, tebakan angka kurang dikit!")
		} else if tebak > a {
			fmt.Print("Salah, tebakan angka terlalu besar!")
		} else if tebak < a {
			fmt.Print("Salah, tebakan angka terlalu kecil!")
		}
	}
	fmt.Print("GG BERHASIL!")

}
