package main

import "fmt"

func main() {
	students := []map[string]string{
		{"name": "Aldi", "score": "A"},
		{"name": "Arif", "score": "B"},
		{"name": "Setiawan", "score": "C"},
	}

	for _, student := range students {
		fmt.Println(student["name"], "->", student["score"])
	}

	Quiz()

	// scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	// var qty int
	// for _, value := range scores {
	// 	qty = qty + value
	// }
	// average := float64(qty) / float64(len(scores))
	// fmt.Println(average)

}
