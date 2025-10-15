package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println("==TABEL PANGKAT 2==")
	for i := 1; i <= n; i++ {
		fmt.Printf("%d %d\n", i, i*i)
	}
}
