package main

import "fmt"

func main() {
	var n, num, total, result int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&num)
		total += num
	}
	result = total / n
	fmt.Println("==HASIL RATA-RATA NILAI==")
	fmt.Print(result)
}
