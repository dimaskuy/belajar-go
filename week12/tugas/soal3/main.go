package main

import "fmt"

func main() {
	var pw string
	success := "kopi123"
	counter := 3

	for pw != success && counter > 0 {
		fmt.Print("Input password: ")
		fmt.Scan(&pw)

		if pw == success {
			fmt.Println("Login sukses lho yhh!🙏🙏🙏")
			break
		}

		counter--
		if counter > 0 {
			fmt.Printf("Salah, coba lagi! (sisa %dx)\n", counter)
		} else {
			fmt.Println("Akses ditolak! hohoho🥀🥀🥀")
		}
	}
}
