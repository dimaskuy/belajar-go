package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&n)

	if n >= 80 && n <= 100 {
		fmt.Println("A")
	} else if n >= 70 {
		fmt.Println("B")
	} else if n >= 60 {
		fmt.Println("C")
	} else if n >= 50 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}
}
