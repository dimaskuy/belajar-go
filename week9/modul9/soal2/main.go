package main

import "fmt"

func main() {
	var siswa, kapasitas, sisaSiswa int // jumlah siswa, kapasitas bus, sisa siswa setelah bus penuh
	var validasi bool                   // variabel untuk cek apakah input valid

	fmt.Scan(&siswa)     // input jumlah siswa
	fmt.Scan(&kapasitas) // input kapasitas bus

	// rentang siswa harus 1-1000, kapasitas harus 1-50
	validasi = (siswa >= 1 && siswa <= 1000) && (kapasitas >= 1 && kapasitas <= 50)
	maxBus := 0  // menyimpan jumlah bus penuh
	nextBus := 0 // jumlah bus tambahan untuk sisa siswa

	// saat input valid
	if validasi {
		// kalkulasi
		maxBus = siswa / kapasitas                        // jumlah bus penuh
		sisaSiswa = siswa % kapasitas                     // sisa siswa yang tidak cukup untuk bus penuh
		nextBus = (sisaSiswa + kapasitas - 1) / kapasitas // jumlah bus tambahan untuk menampung sisa siswa
		fmt.Print(maxBus + nextBus)                       // tampilkan total jumlah bus
	}
	// saat input tidak valid
	if !validasi {
		fmt.Print("Input tak valid, jumlah siswa atau kapasitas kelebihan/kekurangan!") // error handling
	}
}

/*
Program Go ini berfungsi untuk menghitung jumlah bus yang dibutuhkan untuk menampung sejumlah siswa berdasarkan kapasitas tiap bus. Program membaca input ke dalam variabel siswa (jumlah siswa) dan kapasitas (kapasitas per bus), kemudian memvalidasi input menggunakan variabel validasi untuk memastikan siswa berada di antara 1–1000 dan kapasitas di antara 1–50. Jika valid, program menghitung jumlah bus penuh menggunakan variabel maxBus dengan membagi siswa dengan kapasitas, menghitung sisa siswa ke dalam variabel sisaBus, dan menentukan bus tambahan yang dibutuhkan melalui variabel nextBus. Total jumlah bus (maxBus + nextBus) ditampilkan dengan fmt.Print. Jika input tidak valid, program menampilkan pesan kesalahan.
*/
