// 1. Program menerima satu bilangan bulat positif N(N >= 3)
// ii. Dua angka pertama dalam urutan adalah 1 dan 1
// 111. Setiap angka berikutnya dihitung menggunakan rumus: F(i) = (F(i-1)+F(i-2)) mod (i + 2)
// iv. Apabila hasil dari F(i) menghasilkan angka 2 digit (F(i) > 9), maka harus diubah menjadi jumlah dari seluruh digitnya, dilakukan berulang sampai hasilnya satu digit. (ex: F(i) = 13 maka F(i) = 1+3 = 4)

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n && i <= 2; i++ {
		fmt.Print(1, " ")
	}

	start := 1
	accumulator := 1

	for i := 3; i <= n; i++ {
		value := (start + accumulator) % (i + 2)

		for value >= 10 {
			temp := 0
			for j := value; j > 0; j /= 10 {
				temp += j % 10
			}
			value = temp
		}

		fmt.Print(value, " ")
		accumulator = start
		start = value
	}
	fmt.Println("")
}
