package main

import "fmt"

func main() {

	// for i := 1; i < 10; i++ {
	// 	fmt.Println("", i)
	// }

	// i := 1
	// for i <= 100 {
	// 	fmt.Println("", i)
	// 	i++
	// }

	title := "Golang the best language"

	for index, letter := range title {
		if index%2 == 0 {
			fmt.Println("index :", index, " letter :", string(letter))
		}
	}

	for index, letter := range title {
		letterString := string(letter)
		// if letterString == "a" || letterString == "e" || letterString == "i" || letterString == "o" || letterString == "u" {
		// 	fmt.Println("index :", index, " letter :", string(letter))
		// }

		switch letterString {
		case "a", "i", "o", "e", "u":
			fmt.Println("index :", index, " letter :", string(letter))
		}
	}

}
