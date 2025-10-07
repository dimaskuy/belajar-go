package main

import "fmt"

func main() {
	// deklarasi variabel jari-jari (r), tinggi kerucut (tK), dan tinggi tabung (tT)
	var r, tK, tT int
	// assign PI dengan variabel konstant
	const PI = 3.14
	// input variabel & assign variabel tersebut
	fmt.Print("Masukkan jari-jari: ")
	fmt.Scan(&r)
	fmt.Print("Masukkan tinggi kerucut: ")
	fmt.Scan(&tK)
	fmt.Print("Masukkan tinggi tabung: ")
	fmt.Scan(&tT)

	// proses kalkulasi lalu di-assign ke variabel volume kerucut (vK) dan volume tabung (vT) 
	vK := ((1.0/3.0) * PI * float32((r*r) * tK))
	vT := (PI * float32((r * r) * tT))
	// total volume keseluruhan ke variabel vTotal
	vTotal := vK + vT

	// cetak semuanya :)
	fmt.Printf("Volume Tabung: %.2f\nVolume Kerucut: %.2f\nVolume Total: %.2f", vT, vK, vTotal)
}