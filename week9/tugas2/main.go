package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukan: ")
	fmt.Scan(&n)

	if n > 0 && n%2 == 1 {
		fmt.Print("ganjil positif!")
	}
	if n < 0 || n%2 == 0 {
		fmt.Print("bukan bruhhh..")
	}
}
