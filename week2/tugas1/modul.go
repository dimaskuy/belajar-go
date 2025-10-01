package main

import "fmt"


func tugas1() {
    var (
        satu, dua, tiga string
        temp string
    )

    fmt.Print("Masukan input string: ")
    fmt.Scanln(&satu)
    
    fmt.Print("Masukan input string: ")
    fmt.Scanln(&dua)

    fmt.Print("Masukan input string: ")
    fmt.Scanln(&tiga)
    
    fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
    
    temp = satu
    satu = dua
    dua = tiga
    tiga = temp
    
    fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}

func tugas2()  {
    var nama, nim, kelas string
    fmt.Print("Nama: ") 
    fmt.Scan(&nama) 
    fmt.Print("NIM: ") 
    fmt.Scan(&nim) 
    fmt.Print("Kelas: ") 
    fmt.Scan(&kelas)

    fmt.Printf("Halo!\nNama saya %s (NIM %s), dari kelas %s.\nMohon bantuannya!", nama, nim, kelas)
}

func tugas3()  {
    const pi = 3.14159265358979323846
    var r int
    fmt.Print("Masukkan jari-jari: ");
    fmt.Scan(&r);
    hasil := pi*float32(r*r)
    fmt.Printf("\n%.1f", hasil)
}

func tugas4() {
    var f int
	var c float32

	fmt.Print("Masukkan suhu Farenheit: ")
	fmt.Scan(&f)
	c = (5.0/9.0) * float32(f-32)
	fmt.Printf("Suhu dalam Celsius (C) = %.2f", c)
}

func main() {
    // tugas1()
    // tugas2()
    // tugas3()
    // tugas4()
}