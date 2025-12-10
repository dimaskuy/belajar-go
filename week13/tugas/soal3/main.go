package main

import "fmt"

func main() {
	var n float64
	fmt.Scan(&n)

	xLama := n
	var xBaru float64

	for {
		xBaru = 0.5 * (xLama + n/xLama)

		beda := xBaru - xLama
		if beda < 0.0 {
			beda *= -1
		}

		if beda < 0.00001 {
			break
		}

		xLama = xBaru
	}

	fmt.Printf("%.5f\n", xBaru)
}
