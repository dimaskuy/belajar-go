package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	count := 0

	for {
		count++
		n /= 10

		if n == 0 {
			break
		}
	}

	fmt.Println(count)
}
