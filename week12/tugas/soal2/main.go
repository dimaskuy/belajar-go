package main

import "fmt"

func main() {
	n := 1
	total := 0
	for n != 0 {
		fmt.Scan(&n)
		total += n
	}
	fmt.Print("Total: ", total)
}
