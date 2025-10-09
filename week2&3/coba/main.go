package main

import (
	"fmt"
	"strconv"
)

func main() {
	strToInt()
	intToStr()
}

// casting str > int
func strToInt() {
	var teks string = "0"
    tahun, err := strconv.Atoi(teks)
    if err == nil {
		fmt.Println(tahun + 1)
    }
}
// casting int > str
func intToStr() {
    var bil int = 0120312
    var teks string = strconv.Itoa(bil)
    fmt.Println(teks)
}