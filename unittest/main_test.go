package main

import (
	"take-home-test/caseEnam"
	"testing"
)

func TestAddProduct(t *testing.T) {
	cart := &caseEnam.ShoppingCart{Products: make(map[string]caseEnam.Product)}
	cart.AddProduct("p1", "Produk 1", 2)
	if cart.Products["p1"].Kuantitas != 2 {
		t.Error("Diharapkan 2, dapat", cart.Products["p1"].Kuantitas)
	}
	cart.AddProduct("p1", "Produk 1", 3)
	if cart.Products["p1"].Kuantitas != 5 {
		t.Error("Diharapkan 5, dapat", cart.Products["p1"])
	}
}

func TestDeleteProduct(t *testing.T) {
	cart := &caseEnam.ShoppingCart{Products: make(map[string]caseEnam.Product)}
	cart.AddProduct("p1", "Produk 1", 2)
	cart.DeleteProduct("p1")
	if _, ok := cart.Products["p1"]; ok {
		t.Error("Produk yang diharapkan akan dihapus, tetapi masih ada")
	}
	cart.ShowShoppingCart("", 0)
}

func TestShowShoppingCart(t *testing.T) {
	// buat atau tambah baru product
	cart := &caseEnam.ShoppingCart{Products: make(map[string]caseEnam.Product)}
	cart.AddProduct("P001", "Produk 1", 2)
	cart.AddProduct("P002", "Produk 2", 1)
	cart.AddProduct("P003", "Produk 3", 3)
	cart.AddProduct("P003", "Produk 3", 3)

	// Uji coba Filter by nama produk
	filteredProducts := cart.ShowShoppingCart("Produk 2", -1)
	if len(filteredProducts) != 1 {
		t.Errorf("Diharapkan 1 produk, dapat %d", len(filteredProducts))
	}
	if filteredProducts[0].KodeProduk != "P002" {
		t.Errorf("Kode produk yang diharapkan P002, dapat %s", filteredProducts[0].KodeProduk)
	}

	// Uji Coba Filter By Kuantitas
	filteredProducts = cart.ShowShoppingCart("", 1)
	if len(filteredProducts) != 1 {
		t.Errorf("Diharapkan 1 produk, dapat %d", len(filteredProducts))
	}
	if filteredProducts[0].KodeProduk != "P002" {
		t.Errorf("Kode produk yang diharapkan P002, dapat %s", filteredProducts[0].KodeProduk)
	}

	// Uji Coba Filter Tanpa Kuantitas
	filteredProducts = cart.ShowShoppingCart("", -1)
	if len(filteredProducts) != 3 {
		t.Errorf("Diharapkan 3 produk, dapat %d", len(filteredProducts))
	}
}
