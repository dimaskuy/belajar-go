package main

import "fmt"


func tugas1() {
    var x float64
    fmt.Print("Masukkan X: ")
    fmt.Scan(&x)
    hasil := (2/(x-5)) - 5
    fmt.Print("Keluaran: ", hasil)
}

func tugas2()  {
    const pi = 3.14159265358979323846
    var r int
    var v, l float32
    fmt.Print("Jejari: ")
    fmt.Scan(&r)
    v = (4/3)*pi*float32(r*r*r)
    l = 4*pi*float32(r*r)
    fmt.Printf("\nBola dengan jejari %d memiliki volume %.4f dan luas kulit %.4f", r, v, l)
}


func tugas3()  {
    var year int
	fmt.Print("Masukkan tahun: ")
	fmt.Scan(&year)

	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func tugas4() {
    var celsius float64

	fmt.Print("Temperatur Celsius: ")
	fmt.Scan(&celsius)

	reamur := celsius * 4 / 5
	fahrenheit := (celsius * 9 / 5) + 32
	kelvin := celsius + 273

	fmt.Printf("Derajat Reamur: %.2f\n", reamur)
	fmt.Printf("Derajat Fahrenheit: %.2f\n", fahrenheit)
	fmt.Printf("Derajat Kelvin: %.2f\n", kelvin)
}

func main() {
    // tugas1()
    // tugas2()
    // tugas3()
    tugas4()
}