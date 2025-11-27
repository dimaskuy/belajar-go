package main

import "fmt"

func main() {
	var jam24 int
	var jam12 int
	var label string

	fmt.Print("Masukkan jam (0-23): ")
	fmt.Scan(&jam24)

	switch {
	case jam24 == 0:
		jam12 = 12
		label = "AM"
	case jam24 < 12:
		jam12 = jam24
		label = "AM"
	case jam24 == 12:
		jam12 = 12
		label = "PM"
	case jam24 > 12:
		jam12 = jam24 - 12
		label = "PM"
	default:
		jam12 = jam24
		label = "ERROR!"
	}

	if jam24 > 24 {
		fmt.Println("ERROR!")
	} else {
		fmt.Println(jam12, label)
	}
}

// func main() {
// 	var hari string = "2"

// 	switch hari {
// 	case "1", "2":
// 		if hari == "2" {
// 			fmt.Printf("Senin2\n")
// 		}
// 		fmt.Printf("Senin\n")
// 	case "22":
// 		fmt.Printf("Selasa\n")
// 	case "3":
// 		fmt.Printf("Rabu\n")
// 	default:
// 		fmt.Printf("Hari tidak valid\n")
// 	}
// }
