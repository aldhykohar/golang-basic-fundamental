package main

import (
	"fmt"
	"pertama/calculation"
)

func main() {
	fmt.Println("Halo Belajar Golang")

	result := calculation.Add(8, 2)

	fmt.Println(result)
}
