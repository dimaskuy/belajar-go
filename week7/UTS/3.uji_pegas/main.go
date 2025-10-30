/* Besar gaya pegas berbanding lurus dengan pertambahan panjangnya, dengan rumus: */
// F = κ*μ
// Di mana:
// F = Gaya pegas (Newton)
// K = Konstanta Pegas (N/m)
// μ = Pertambahan Panjang Pegas (m)

/* PERHITUNGAN */
// 1. Gaya setiap Pegas menggunakan rumus
// F = κχμ
// 2. Total Gaya seluruh Pegas (Penjumlahan Semua F₁) Ftotal = Σ i=1 * Fi
// 3. Rata rata pertambahan panjang semua pegas = Σn i = 1 * μi / n
// 4. Menentukan apakah sistem pegas kuat atau tidak, dengan logika sebagai berikut:
// a. Sistem kuat jika total gaya ≥ 100 Newton
// b. Sistem tidak kuat jika total gaya < 100 Newton
// Kuat Ftotal ≥ 100

/* Input Format */
// Int n = jumlah pegas
// Fi  , к  , μ, Ϝ_total, μ_rata bertipe data float32
// Kuat bertipe data Boolean
// Untuk setiap iterasi ke-i, akan menjalankan:
// Input к  dan  μ
// Perhitugnan Fi

/* Constraints */
// 1. Tidak boleh menggunakan percabangan (If, else, Switch)
// 2. Fi  boleh di outputkan pada setiap iterasi untuk input к  dan μ
// 3. Fi  , к  , μ bertipe data float32
// 4. Hasil ditampilkan dengan format 2 angka dibelakang koma

/* Output Format */
// 1. Menampilkan Gaya (Fi ) pada setiap Iterasi
// 2. Menampilkan total gaya seluruh pegas
// 3. Tampilkan rata rata pertambahan panjang
// 4. Tampilkan status kuat (True/False) sesuai kondisi diatas

package main

import "fmt"

func main() {
	var n int
	var besarF, constant, pPegas, fTotal, pTotal, pRata float32
	var isStronk bool

	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&constant)
		fmt.Scan(&pPegas)
		besarF = constant * pPegas
		fmt.Printf("Pegas ke-%d menghasilkan gaya: %.2f N\n", i, besarF)
		fTotal += besarF
		pTotal += pPegas
	}

	pRata = pTotal / float32(n)
	isStronk = fTotal >= 100

	fmt.Println("----------------------------------------")
	fmt.Println("Data Uji Pegas Paman Berkacamata")
	fmt.Println("----------------------------------------")
	fmt.Printf("Total gaya seluruh pegas: %.2f N\n", fTotal)
	fmt.Printf("Rata-rata pertambahan panjang: %.2f m\n", pRata)
	fmt.Printf("Apakah sistem pegas kuat (>=100N)? %t\n", isStronk)
	fmt.Println("----------------------------------------")
}
