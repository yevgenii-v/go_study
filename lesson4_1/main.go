package main

import (
	"fmt"
	"lessong4_1/calculator"
	"lessong4_1/greeting"
)

func main() {
	greeting.Hello()

	fmt.Printf("Сума: %d\n", calculator.Add(4, 5))
	fmt.Printf("Різниця: %d\n", calculator.Subtract(13, 2))
}
