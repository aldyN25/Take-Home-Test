package main

import (
	"fmt"
	caseDua "take-home-test/caseDua"
	"take-home-test/caseEnam"
	caseSatu "take-home-test/caseSatu"
)

func main() {

	fmt.Println("******** JAWABAN SOAL 1 ********")
	harga1 := 145000
	harga2 := 2099

	fmt.Println(caseSatu.Pecahan(harga1))
	fmt.Println(caseSatu.Pecahan(harga2))

	fmt.Println("")

	fmt.Println("******** JAWABAN SOAL 2 ********")
	fmt.Println(caseDua.OneEditRule("telkom", "telecom"))
	fmt.Println(caseDua.OneEditRule("telkom", "tlkom"))

	fmt.Println("")
	fmt.Println("******* JAWABAN SOAL 6 *******")
	ShoppingCart := caseEnam.NewShoppingCart()
	ShoppingCart.AddProduct("P001", "Produk 1", 2)
	ShoppingCart.AddProduct("P002", "Produk 2", 1)
	ShoppingCart.AddProduct("P003", "Produk 3", 3)
	ShoppingCart.AddProduct("P003", "Produk 3", 3)

	fmt.Println("Daftar ShoppingCart:")
	for _, product := range ShoppingCart.ShowShoppingCart("", 0) {
		fmt.Printf("%s - %s - (%d)\n", product.KodeProduk, product.NamaProduk, product.Kuantitas)
	}

	fmt.Println("Filter ShoppingCart (NamaProduk = Produk 3):")
	for _, product := range ShoppingCart.ShowShoppingCart("Produk 3", 0) {
		fmt.Printf("%s - %s - (%d)\n", product.KodeProduk, product.NamaProduk, product.Kuantitas)
	}

	// Hapus Data Produk dengan Kode P002
	ShoppingCart.DeleteProduct("P002")

	fmt.Println()
	fmt.Println("Daftar ShoppingCart Setelah Dihapus:")

	for _, product := range ShoppingCart.ShowShoppingCart("", 0) {
		fmt.Printf("%s - %s - (%d)\n", product.KodeProduk, product.NamaProduk, product.Kuantitas)
	}

}
