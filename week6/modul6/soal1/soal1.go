package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n) // init variabel n & simpan nilai nya

	// looping dari n ke 1 secara menurun (decrement)
	for i := n; i > 0; i-- {
		fmt.Printf("%d ", i) // cetak nilai i-nya dan diikuti spasi
	}
}
