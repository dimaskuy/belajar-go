package main

import "fmt"

func main() {
	var tipe int
	var n int

	fmt.Print("Masukkan tipe: ")
	fmt.Scan(&tipe)
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)

	switch tipe {
	// kiri
	case 1:
		for i := 0; i <= n; i++ {
			for j := 0; j < i; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	// tengah
	case 2:
		var item string
		for i := 0; i < 10; i++ {
			for i := 0; i <= n; i++ {
				for j := 0; j < i; j++ {
					fmt.Print("*")
				}
				fmt.Println()
			}
			for j := 10; j >= i; j-- {
				item += " "
				// fmt.Print("")
			}
			for k := 1; k <= i; k++ {
				item += "*"
				// fmt.Print("*")
			}
			for l := 0; l <= i; l++ {
				item += "*"
				// fmt.Print("*")
			}
			// item += "\n"
			fmt.Println()
		}
		fmt.Print(item)
	// kanan
	case 3:
		for i := n; i <= n; i-- {
			for j := 0; j <= i; j++ {
				fmt.Print("*")
			}
			fmt.Println("")
		}

	default:
		fmt.Print("ERROR!")
	}
}
