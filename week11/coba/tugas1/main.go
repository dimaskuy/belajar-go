package main

import "fmt"

func main() {
	var n int = 90
	// var sisa

	switch {
	case n <= 100 && n >= 90:
		fmt.Println('A')
		if n == 97 {
			fmt.Println("Outstanding!")
		}
		// if n % 2 != 0 && n % 2 {
		// 	fmt.Println("Outstanding!")
		// }
	case n <= 89 && n >= 80:
		fmt.Println("B")
	case n <= 79 && n >= 70:
		fmt.Println("C")
	case n <= 69 && n >= 60:
		fmt.Print("D")
	case n <= 59 && n >= 0:
		fmt.Print("E")
		if n%2 == 0 {
			fmt.Println("Retake")
		}
	default:
		fmt.Print("ERROR JIR")
	}
}
