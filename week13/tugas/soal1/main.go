package main

import "fmt"

func main() {
	var n int
	i := 0
	total := 0.0

	for {
		fmt.Scan(&n)
		if n == -1 {
			break
		}
		if n <= 0 {
			continue
		}
		i++
		total += float64(n)
	}

	fmt.Printf("%d %.2f", i, total/float64(i))
}
