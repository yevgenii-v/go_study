package main

import "fmt"

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func (p *Product) PrintProduct() {
	fmt.Println("Назва:", p.Name)
	fmt.Println("Ціна:", p.Price)
	fmt.Println("Кількість:", p.Quantity)
}

func (p *Product) IsAvailable() bool {
	return p.Quantity > 0
}

func (p *Product) ChangePrice(newPrice float64) {
	p.Price = newPrice
}
