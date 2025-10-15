package main

import "fmt"

// for (inisialisasi); (kondisi); (update/decrement/increment)

func main() {
	coba2()
}

func coba1() {
	for i := 1; i <= 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Print("JAWA ")
		}
		fmt.Println(" ")
	}
}

func coba2() {
	var iterasi, n int
	n = 1000

	for iterasi = 1; iterasi <= n; iterasi++ {
		fmt.Println(iterasi, "CAK1BAB3 Algoritma Pemrograman 1")
	}
}

func coba3() {

}
