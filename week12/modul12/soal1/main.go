package main

import "fmt"

func main() {
	var user, pw string
	failCount := 0

	for {
		fmt.Scan(&user, &pw)

		if user == "Admin" && pw == "Admin" {
			fmt.Printf("%d percobaan gagal login\n", failCount)
			break
		}

		failCount++
	}
}
