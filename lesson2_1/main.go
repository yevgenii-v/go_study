package main

import (
	"fmt"
	"time"
)

func main() {
	hello()
	fmt.Println(summary(6, 8))
	fmt.Println(checkDate())
}

func hello() {
	fmt.Print("Hello World!\n")
}

func summary(a, b int) int {
	return a + b
}

func checkDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
