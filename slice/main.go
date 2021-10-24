package main

import "fmt"

func main() {
	var gamingConsoles []string

	gamingConsoles = append(gamingConsoles, "Playstation4")
	gamingConsoles = append(gamingConsoles, "Nintendo")
	gamingConsoles = append(gamingConsoles, "Xbox")

	// gamingConsoles := []string{"Playstation4", "Nintendo", "Xbox"}

	fmt.Println(gamingConsoles)
}
