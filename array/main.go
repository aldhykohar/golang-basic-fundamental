package main

import "fmt"

func main() {
	// var languages [5]string
	// languages[0] = "GO"
	// languages[1] = "Java"
	// languages[2] = "Kotlin"
	// languages[3] = "C#"
	// languages[4] = "Python"

	// languages := [5]string{
	// 	"GO",
	// 	"Java",
	// 	"Kotlin",
	// 	"C#",
	// 	"Python",
	// }

	languages := [...]string{
		"GO",
		"Java",
		"Kotlin",
		"C#",
		"Python",
		"Javascript",
	}

	fmt.Println(languages)

	length := len(languages)
	fmt.Println(length)

	for _, lang := range languages {
		fmt.Println("value :", lang)
	}
}
