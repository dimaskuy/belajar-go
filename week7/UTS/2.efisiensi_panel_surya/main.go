// Diberikan rumus daya listrik yang dihasilkan oleh satu panel:
// P = Ι·Α·η(0)
// Dengan:
// P = daya listrik (Watt)
// I = intensitas cahaya matahari (W/m²), diasumsikan konstan = 1000 W/m²
// A = luas panel = 1.5 m²
// η(θ) = efisiensi tergantung sudut, dihitung sebagai: eta(theta) = c * os((theta*pi)/180)
// (konversi derajat ke radian, lalu ambil nilai cosinus)

/* Input Format */
// 1.	Bilangan bulat N (jumlah panel surya yang diuji, 1≤ N ≤ 100 )
// 2.	Untuk setiap panel ke-i (dari 1 sampai N):
//         • Sudut pemasangan dalam derajat (bilangan bulat, ) Sudut pemasangan θi dalam derajat (bilangan bulat, 0 ≤ θi ≤ 90 )

/* Constraints */
// •	1 ≤ N ≤ 100
// •	1 ≤ mi  ≤ 1000 (bilangan bulat)
// •	0 ≤ vi ≤500 (bilangan riil)
// •	Tidak boleh menggunakan if, else, switch, atau operator ternary
// •	Hanya boleh menggunakan for-loop, I/O, variabel, casting, dan operasi aritmetika

/* Output Format */
// 1.	Untuk setiap panel, tampilkan
//     Panel ke-i: θ=XX°, Daya=YY.YY W
//     (format daya dengan 2 angka di belakang koma)
// 2.	Setelah semua panel ditampilkan, tampilkan:
//         •	Total daya seluruh panel (dalam Watt, 2 angka di belakang koma)
//         •	Rata-rata efisiensi sistem (dalam desimal, 4 angka di belakang koma)
//         •	Jumlah panel yang efisiensinya ≥ 0.8 (hitung berapa panel dengan θ ≤ 18 derajat)

package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var totalsDaya, totalsSudut float64
	var jmlhPanelEff float64

	for i := 1; i <= n; i++ {
		var teta float64
		fmt.Scan(&teta)

		watt := 1500 * (1 - teta/90)

		fmt.Printf("Panel ke-%d: sudut=%.0f derajat, Daya=%.2f W\n", i, teta, watt)

		totalsDaya += watt
		totalsSudut += teta

		jmlhPanelEff += float64(int((19 - teta) / 18.5))
	}

	rataEff := totalsDaya / (float64(n) * 1500)
	rataSudut := totalsSudut / float64(n)

	fmt.Println("----------------------------------------")
	fmt.Printf("Total daya: %.2f W\n", totalsDaya)
	fmt.Printf("Rata-rata efisiensi: %.4f\n", rataEff)
	fmt.Printf("Rata-rata sudut: %.2f derajat\n", rataSudut)
}
