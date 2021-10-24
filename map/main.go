package main

import "fmt"

func main() {

	//can use this is to

	// var myMap map[string]int
	// myMap = map[string]int{}

	// myMap["Ruby"] = 9
	// myMap["Go"] = 9
	// myMap["Js"] = 8
	// myMap["Go"] = 10

	myMap := map[string]string{
		"ruby": "is beauty",
		"go":   "is fast",
		"Js":   "hype",
	}

	// fmt.Println(myMap)

	for key, value := range myMap {
		fmt.Println("Key :", key, " Value :", value)
	}

	fmt.Println("==========")

	//delete value in map
	delete(myMap, "ruby")

	//check if key is available in map
	values, isAvaiable := myMap["python"]
	fmt.Println(isAvaiable)
	fmt.Println(values)
}
