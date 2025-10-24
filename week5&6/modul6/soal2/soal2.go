package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n) // init variabel n & simpan nilai nya

	// looping dari n ke 1 secara menurun (decrement)
	for i := n; i > 0; i-- {
		for j := 1; j <= i; j++ { // cetak dari 1 hingga i
			fmt.Printf("%d ", j)
		}
		fmt.Println(" ") // ganti ke baris baru
	}
}
