package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	digit := 0
	for i := num; i > 0; i /= 10 {
		digit++
	}
	fmt.Println("==BANYAK DIGIT==")
	fmt.Print(digit)
}
