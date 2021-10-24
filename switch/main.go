package main

import "fmt"

func main() {
	number := 5

	switch number {
	case 1:
		fmt.Println("Satu")
	case 2:
		fmt.Println("Dua")
	default:
		fmt.Println("DEFAULT")
	}
}
