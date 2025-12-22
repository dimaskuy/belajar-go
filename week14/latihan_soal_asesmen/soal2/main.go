package main

import "fmt"

// REPEAT-UNTIL LOOP
func main() {
	var awal, akhir int
	fmt.Scan(&awal)
	fmt.Scan(&akhir)

	if awal%2 == 1 {
		fmt.Print(awal, " ")
	}

	for {
		awal++

		if awal%2 == 1 {
			fmt.Print(awal, " ")
		}

		if awal >= akhir {
			break
		}
	}
}

// ALTERNATIF: (FOR LOOP)
// func main() {
// 	var awal, akhir int
// 	fmt.Scanln(&awal)
// 	fmt.Scanln(&akhir)

// 	fmt.Print(awal, " ")
// 	for i := awal; i < akhir; i = i + 2 {
// 		fmt.Print(i, " ")
// 	}
// }
