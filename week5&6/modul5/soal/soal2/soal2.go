package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n) // input jumlah bilangan/bulan
	a := 0       // nilai awal deret
	b := 1       // nilai awal kedua deret
	temp := a    // variabel bantu untuk menyimpan nilai a sementara

	for i := 1; i <= n; i++ { // loop dari 1 hingga n
		fmt.Print(a, " ")
		temp = a     // simpan nilai a ke variabel temp
		a = b        // geser nilai b menjadi nilai a berikutnya
		b = temp + b // update nilai b (a + b lama) utk bilangan deret selanjutnya
	}
	//print: 0 1 1 2 3 5
	// temp: 0 1 1 2 3
	// a: 1 1 2 3 5
	// b: 1 2 3 5 8

}
