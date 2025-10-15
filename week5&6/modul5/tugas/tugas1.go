package main

import "fmt"

func main() {
	var n, result int
	fmt.Scan(&n)
	for i := 1; i <= n; i += 2 {
		result += i
	}
	fmt.Println("==DERET BILANGAN GANJIL==")
	fmt.Print(result)
}
