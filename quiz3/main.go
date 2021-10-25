package main

import (
	"errors"
	"fmt"
)

func main() {
	array := []int{10, 5, 8, 9, 7}
	sum := sum(array)
	fmt.Println(sum)

	result, err := calculate(4, 2, "d")
	if err != nil {
		fmt.Println("Terjadi kesalahan :", err.Error())
	}
	fmt.Println(result)
}

func sum(array []int) (total int) {
	for _, value := range array {
		total = total + value
	}
	return total
}

func calculate(value1, value2 int, operator string) (result int, err error) {
	switch operator {
	case "+":
		result = value1 + value2
	case "-":
		result = value1 - value2
	case "*":
		result = value1 * value2
	case "/":
		result = value1 / value2
	default:
		err = errors.New("Unknown operation")
	}
	return
}
