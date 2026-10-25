package main

import "fmt"

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func main() {
	product := Product{
		Name:     "Телефон",
		Price:    14000.50,
		Quantity: 2,
	}

	printProduct(product)

	product.Name = "Тостер"
	product.Price = 1500
	printProduct(product)
}

func printProduct(p Product) {
	fmt.Println("Назва:", p.Name)
	fmt.Println("Ціна:", p.Price)
	fmt.Println("Кількість:", p.Quantity)
	fmt.Print("\n")
}
