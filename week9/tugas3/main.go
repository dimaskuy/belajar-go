package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukan X dan Y: ")
	fmt.Scan(&x, &y)

	if x%y == 0 {
		fmt.Println(true)
	}
	if x%y != 0 {
		fmt.Println(false)
	}
	if y%x == 0 {
		fmt.Println(true)
	}
	if y%x != 0 {
		fmt.Println(false)
	}
}
