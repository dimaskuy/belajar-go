package main

import "fmt"

func main() {
	var berat int
	var tinggi float32
	fmt.Print("Masukan berat (kg): ")
	fmt.Scan(&berat)
	fmt.Print("Masukan tinggi (m): ")
	fmt.Scan(&tinggi)
	bmi := float32(berat) / (tinggi * tinggi)

	fmt.Printf("BMI Anda: %.2f ", bmi)
	if bmi < 18.5 && bmi > 0 {
		fmt.Print("(Kurus)")
	} else if bmi < 25 {
		fmt.Print("(Normal)")
	} else if bmi < 30 {
		fmt.Print("(Overweight)")
	} else if bmi < 35 {
		fmt.Print("(Obesitas tingkat 1)")
	} else if bmi < 40 {
		fmt.Print("(Obesitas tingkat 2)")
	} else if bmi >= 40 {
		fmt.Print("(Obesitas tingkat 3)")
	} else {
		fmt.Print(", yakin inputnya bener??!")
	}
}
