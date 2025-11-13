package main

import "fmt"

func main() {
	var jmlh_orang int
	fmt.Print("Masukan: ")
	fmt.Scan(&jmlh_orang)
	var mobil int

	if jmlh_orang%4 == 0 {
		mobil = (jmlh_orang / 4)
		fmt.Print(mobil)
	}
	if jmlh_orang > 4 && jmlh_orang%4 != 0 {
		mobil = (jmlh_orang / 4) + 1
		fmt.Print(mobil)
	}
}
