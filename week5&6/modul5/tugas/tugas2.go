package main

import "fmt"

func main() {
	var n, harga_satuan, jumlah_unit, result int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&harga_satuan, &jumlah_unit)
		result += harga_satuan * jumlah_unit
	}
	fmt.Println("==TOTAL BIAYA BELANJA==")
	fmt.Print(result)
}
