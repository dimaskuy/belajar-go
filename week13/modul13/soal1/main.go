package main

import "fmt"

func main() {
	var n int
	tebakan := 0
	fmt.Print("Rahasia Riri (1-10): ")
	fmt.Scanln(&n)

	if n > 10 || n <= 0 {
		fmt.Print("Tidak valid!!!")
	} else {
		for {
			fmt.Print("Roro Menjawab: ")
			fmt.Scanln(&tebakan)
			if tebakan != n {
				if tebakan == n-1 || tebakan == n-2 || tebakan == n+1 || tebakan == n+2 {
					fmt.Println("Riri Menjawab: Kamu Salah! (Dikit lagi...)")
				} else {
					fmt.Println("Riri Menjawab: Kamu Salah!!!")
				}
			} else {
				fmt.Print("Riri Menjawab: Kamu Benar bung!!!")
				break
			}
		}
	}
}
