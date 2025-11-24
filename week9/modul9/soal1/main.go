package main

import "fmt"

func main() {
	var n int
	var counter int = 0 // penghitung faktor prima
	var prima bool      // pendanda apakah suatu angka adalah prima
	var faktor bool     // pendanda apakah suatu angka adalah faktor n
	fmt.Scan(&n)        // sebuah input dari variabel n

	for i := 1; i <= n; i++ { // pencari faktor dari n
		faktor = n%i == 0 // faktor akan true jika n modulus 1
		prima = i > 1     // bilangan prima harus > 1

		for j := 2; j*j <= i && prima; j++ { // pencari i adalah prima
			if i%j == 0 { // jika i habis dibagi j, maka i bukan prima
				prima = false
			}
		}

		if faktor && prima { // jika i adalah faktor dan bilangan prima
			counter += 1 // tandai jumlah faktor prima dengan menambahkan 1
		}
	}

	fmt.Println(counter) // cetak jumlah faktor prima
}

/*
Program Go di atas berfungsi untuk menghitung jumlah faktor prima dari suatu bilangan n yang dimasukkan oleh pengguna. Program dimulai dengan membaca input integer n melalui fmt.Scan(&n). Lalu, dengan perulangan for i := 1; i <= n; i++, program mencari semua faktor dari n (yaitu nilai i yang membagi n tanpa sisa). Untuk setiap faktor tersebut, program juga mengecek apakah i merupakan bilangan prima — dilakukan dengan memeriksa apakah i memiliki pembagi lain selain 1 dan dirinya sendiri. Jika i adalah faktor dan sekaligus bilangan prima, maka variabel counter ditambah 1. Pada akhir program, jumlah faktor prima tersebut ditampilkan dengan fmt.Println(counter).
*/
