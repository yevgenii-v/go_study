package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const SUM = "+"
const SUB = "-"
const MULTI = "*"
const DIV = "/"

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter first number: ")
	inp1, err1 := reader.ReadString('\n')
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	inp1 = strings.TrimSpace(inp1)
	if !canBeInt(inp1) {
		fmt.Println("Not an integer or out of range")
		return
	}

	fmt.Print("Enter second number: ")
	inp2, err2 := reader.ReadString('\n')
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	inp2 = strings.TrimSpace(inp2)
	if !canBeInt(inp2) {
		fmt.Println("Not an integer or out of range")
		return
	}

	fmt.Print("Enter operation: ")
	oper, err3 := reader.ReadString('\n')
	oper = strings.TrimSpace(oper)
	if err3 != nil {
		fmt.Println(err3)
		return
	}

	if oper == DIV && inp2 == "0" {
		fmt.Println("Divide by 0 forbidden")
		return
	}

	switch oper {
	case SUM:
		fmt.Println(sum(inp1, inp2))
	case SUB:
		fmt.Println(sub(inp1, inp2))
	case MULTI:
		fmt.Println(mul(inp1, inp2))
	case DIV:
		fmt.Println(div(inp1, inp2))
	default:
		fmt.Println("Unknown operation")
	}
}

func sum(arg1, arg2 string) int64 {
	return convStrToInt(arg1) + convStrToInt(arg2)
}

func sub(arg1, arg2 string) int64 {
	return convStrToInt(arg1) - convStrToInt(arg2)
}

func mul(arg1, arg2 string) int64 {
	return convStrToInt(arg1) * convStrToInt(arg2)
}

func div(arg1, arg2 string) float64 {
	return convStrToFloat64(arg1) / convStrToFloat64(arg2)
}

func convStrToInt(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	return i
}

func convStrToFloat64(s string) float64 {
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println(err)
	}

	return i
}

func canBeInt(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return false
	}

	return true
}
