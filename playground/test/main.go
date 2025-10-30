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
