package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	awal := n    // angka awal sebelum di-kalkulasi
	reverse := 0 // variabel yg akan dibalikkan angka-nya

	for i := n; i > 0; i = i / 10 {
		reverse = (reverse * 10) + (n % 10) // mengambil digit terakhir
		n /= 10                             // menghilangkan digit terakhir hingga habis
	}

	fmt.Println(awal == reverse) // jika angka palindrome output "true", jika tak sama output "false"
}
