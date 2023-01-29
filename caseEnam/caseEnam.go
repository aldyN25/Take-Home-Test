package caseEnam

import (
	"sync"
)

type Product struct {
	KodeProduk string
	NamaProduk string
	Kuantitas  int
}

type ShoppingCart struct {
	Products map[string]Product
	mutex    sync.RWMutex
}

func NewShoppingCart() *ShoppingCart {
	return &ShoppingCart{
		Products: make(map[string]Product),
	}
}

func (c *ShoppingCart) AddProduct(KodeProduk, NamaProduk string, Kuantitas int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	product, ok := c.Products[KodeProduk]
	if ok {
		product.Kuantitas += Kuantitas
	} else {
		product = Product{KodeProduk: KodeProduk, NamaProduk: NamaProduk, Kuantitas: Kuantitas}
	}
	c.Products[KodeProduk] = product
}

func (c *ShoppingCart) DeleteProduct(KodeProduk string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.Products, KodeProduk)
}

func (c *ShoppingCart) ShowShoppingCart(filterName string, filterKuantitas int) []Product {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	result := []Product{}

	for _, product := range c.Products {
		if (filterName == "" || product.NamaProduk == filterName) &&
			(filterKuantitas <= 0 || product.Kuantitas == filterKuantitas) {
			result = append(result, product)
		}
	}
	return result
}
