package main

import "fmt"

// mendefinisikan tipe bentukan (struct) "Karyawan"
type Karyawan struct {
	nama                   string
	gajiPokok, persenPajak float32
}

func main() {
	// memanggil struct yg sudah didefinisikan
	var karyawan Karyawan
	// input
	fmt.Print("Nama karyawan: ")
	fmt.Scan(&karyawan.nama)
	fmt.Print("Gaji pokok: Rp ")
	fmt.Scan(&karyawan.gajiPokok)
	fmt.Print("Persen pajak (%): ")
	fmt.Scan(&karyawan.persenPajak)

	// kalkulasi pajak dan gaji bersih
	pajakFinal := (karyawan.persenPajak / 100) * karyawan.gajiPokok
	gajiBersih := karyawan.gajiPokok - pajakFinal

	// output hasil
	fmt.Println("\n=== Slip Gaji Karyawan ===")
	fmt.Println("Nama Karyawan:", karyawan.nama)
	fmt.Printf("Gaji Pokok: %.0f\n", karyawan.gajiPokok)
	fmt.Printf("Pajak: %.0f\n", pajakFinal)
	fmt.Printf("Gaji Bersih: %.0f\n", gajiBersih)
}
