package casesatu

import (
	"fmt"
	"math"
)

func Pecahan(harga int) map[string]int {
	pecahan := make(map[string]int)
	data := []int{100000, 20000, 5000, 2000, 1000, 500, 100}

	var jumlah int
	for _, val := range data {
		jumlah = harga / val
		// fmt.Println(harga)
		if harga > 0 && harga < 100 {
			harga = 100
		}
		if jumlah > 0 {
			pecahan[fmt.Sprintf("Rp. %d", val)] = jumlah
			// fmt.Println("harga : ", harga)
			// fmt.Println("jumlah :", pecahan)
			harga -= int(math.Ceil(float64(jumlah) * float64(val)))
		}
	}

	return pecahan
}
