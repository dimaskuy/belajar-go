package main

import "fmt"

// func main() {
// 	var n, massa, berat int
// 	var result float32
// 	fmt.Scan(&n)
// 	for i := 1; i <= n; i++ {
// 		fmt.Scan(&massa, &berat)
// 		result = result + float32(massa*berat)
// 	}
// 	result /= 3

// 	fmt.Print(result)
// }

func main() {
	var n int
	var op, result bool
	result = true
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&op)
		result = op && result
	}
	fmt.Print(result)
}
