package main

import "fmt"

func main() {
	var stamina int
	fmt.Scanln(&stamina)

	if stamina == -99 {
		fmt.Println(0, " ", 0)
	} else {
		min := stamina
		max := stamina

		for {
			fmt.Scanln(&stamina)
			if stamina == -99 {
				break
			}

			if stamina >= max {
				max = stamina
			}
			if stamina <= min {
				min = stamina
			}
		}
		fmt.Println(max, " ", min)
	}
}

// ALTERNATIF
// func main() {
// 	var stamina int
// 	fmt.Scanln(&stamina)
// 	first := true

// 	for {
// 		fmt.Scanln(&stamina)

// 		if first {
// 			min := stamina
// 			max := stamina
// 			first = false
// 		}

// 		if stamina == -99 {
// 			break
// 		}

// 		if stamina >= max {
// 			max = stamina
// 		}
// 		if stamina <= min {
// 			min = stamina
// 		}
// 	}
// 		fmt.Println(max, " ", min)
// }
