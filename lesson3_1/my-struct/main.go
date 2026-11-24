package main

import "fmt"

func main() {
	p1 := Product{
		Name:     "Телефон",
		Price:    14000.50,
		Quantity: 0,
	}

	p1.PrintProduct()
	fmt.Println("В наявності:", p1.IsAvailable())
	fmt.Print("\n")

	p2 := Product{
		Name:     "Тостер",
		Price:    1500,
		Quantity: 2,
	}

	p2.PrintProduct()
	fmt.Println("В наявності:", p2.IsAvailable())
	fmt.Print("\n")

	fmt.Println("Тостер отримав знижку.")
	fmt.Print("\n")
	p2.ChangePrice(1200)
	p2.PrintProduct()
}
