package service

type IProductInterface interface {
	AddProduct(string, string, float32) int
	UpdateProduct(int, string, string, float32)
	DeleteProduct(int)
}

type Product struct {
	IProduct IProductInterface
}

func (p *Product) Add() int {
	p.IProduct.AddProduct("bolgarka", "stihl", 1500.5)

	return 1
}

func (p *Product) Update() {
	p.IProduct.UpdateProduct(1, "usm", "stihl", 3000.8)
}

func (p *Product) Delete() {
	p.IProduct.DeleteProduct(1)
}
