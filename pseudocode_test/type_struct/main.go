package main

import (
	"fmt"
	"math"
)

// kamus data
const (
	PAJAK           = 0.10
	BIAYA_TRANSAKSI = 0.002
)

type Investasi struct {
	hargaBeli          float64
	hargaJual          float64
	jumlahSaham        int
	totalInvestasiAwal float64
	totalPenjualan     float64
	keuntunganKotor    float64
	biayaTransaksi     float64
	pajakKeuntungan    float64
	keuntunganBersih   float64
}

// algoritma
func main() {
	var saham Investasi

	// Input
	fmt.Print("Masukkan harga beli: Rp ")
	fmt.Scanln(&saham.hargaBeli)

	fmt.Print("Masukkan harga jual: Rp ")
	fmt.Scanln(&saham.hargaJual)

	fmt.Print("Masukkan jumlah saham: Rp ")
	fmt.Scanln(&saham.jumlahSaham)

	// Proses
	saham.totalInvestasiAwal = saham.hargaBeli * float64(saham.jumlahSaham)
	saham.totalPenjualan = saham.hargaJual * float64(saham.jumlahSaham)
	saham.keuntunganKotor = saham.totalPenjualan - saham.totalInvestasiAwal
	saham.biayaTransaksi = BIAYA_TRANSAKSI * saham.totalPenjualan
	saham.pajakKeuntungan = PAJAK * math.Max(0, saham.keuntunganKotor)
	saham.keuntunganBersih = saham.keuntunganKotor - saham.biayaTransaksi - saham.pajakKeuntungan

	// Output
	fmt.Println("=========================================")
	fmt.Printf("Total Investasi Awal : Rp %.2f\n", saham.totalInvestasiAwal)
	fmt.Printf("Total Penjualan      : Rp %.2f\n", saham.totalPenjualan)
	fmt.Printf("Keuntungan Kotor     : Rp %.2f\n", saham.keuntunganKotor)
	fmt.Printf("Biaya Transaksi      : Rp %.2f\n", saham.biayaTransaksi)
	fmt.Printf("Pajak Keuntungan     : Rp %.2f\n", saham.pajakKeuntungan)
	fmt.Printf("Keuntungan Bersih    : Rp %.2f\n", saham.keuntunganBersih)
	fmt.Println("=========================================")
}
