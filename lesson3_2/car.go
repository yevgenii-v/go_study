package main

type Car struct {
	Brand string
	Model string
	Year  int
}

func (c *Car) PrintInfo() {
	print(
		"Brand: ", c.Brand, "\n",
		"Model: ", c.Model, "\n",
		"Рік: ", c.Year, "\n",
	)
}

func (c *Car) IsOld(currentYear int) bool {
	return currentYear-c.Year > 20
}
