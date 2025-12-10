package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for n != 0 {
		hasil := n % 10
		fmt.Println(hasil)
		n /= 10
	}
}
