package main

import (
	"fmt"
	"strconv"
)

func main() {
	var tipe string
	var persenDiskon, tarif, diskon, jarak, result int

	fmt.Print("Masukkan jenis kendaraan (Mobil/Bus/Truk): ")
	fmt.Scan(&tipe)
	fmt.Print("Masukkan jarak tempuh (km): ")
	fmt.Scan(&jarak)

	if tipe == "Mobil" || tipe == "mobil" || tipe == "MOBIL" {
		tarif = jarak * 1000
	} else if tipe == "Bus" || tipe == "bus" || tipe == "BUS" {
		tarif = jarak * 1500
	} else if tipe == "Truk" || tipe == "truk" || tipe == "TRUK" {
		tarif = jarak * 2000
	} else {
		fmt.Println("KENDARAAN TIDAK VALID!")
	}

	if jarak < 20 && jarak > 0 {
		persenDiskon = 0
	} else if jarak < 50 {
		persenDiskon = 5
	} else if jarak < 100 {
		persenDiskon = 15
	} else {
		persenDiskon = 25
	}

	diskon = int(float32(tarif) * (float32(persenDiskon) / 100))
	result = tarif - diskon
	persen := strconv.Itoa(persenDiskon) + "%"

	fmt.Printf("Jenis Kendaraan: %s\nJarak Tempuh: %d km\nTarif Normal: Rp %d\nDiskon: %s\nTotal Bayar: Rp %d", tipe, jarak, tarif, persen, result)
}
