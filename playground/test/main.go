package main

import "fmt"

func main() {
	var v, t, awal int
	fmt.Scan(&awal)
	fmt.Scan(&v)
	fmt.Scan(&t)

	jarak := v * t
	result := jarak + awal

	fmt.Print(result)
}
