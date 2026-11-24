package main

func main() {
	car := Car{
		"Lamborghini",
		"Murcielago",
		2001,
	}

	car.PrintInfo()
	println("Стара модель:", car.IsOld(2026), "\n")

	car2 := Car{
		"Lamborghini",
		"Aventador",
		2023,
	}

	car2.PrintInfo()
	println("Стара модель:", car2.IsOld(2026))
}
