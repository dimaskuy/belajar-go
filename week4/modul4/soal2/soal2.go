package main

import "fmt"

func main() {
	// deklarasi variabel terkait
	var uts, uas, tugas float32
	// inisiasi bobot nilai
	const bobotUTS = 0.30
	const bobotUAS = 0.40
	const bobotTugas = 0.30

	// input
	fmt.Print("Masukkan nilai UTS: ")
	fmt.Scan(&uts)
	fmt.Print("Masukkan nilai UAS: ")
	fmt.Scan(&uas)
	fmt.Print("Masukkan nilai Tugas: ")
	fmt.Scan(&tugas)
	// kalkulasi nilai akhir
	nilaiFinal := (uts * bobotUTS) + (uas * bobotUAS) + (tugas * bobotTugas)

	// output
	fmt.Printf("Nilai Akhir Mahasiswa: %.1f", nilaiFinal)
}
