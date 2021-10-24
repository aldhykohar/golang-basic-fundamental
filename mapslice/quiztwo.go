package main

import "fmt"

func Quiz() {
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	var goodScore []int

	var qty int
	for _, value := range scores {
		qty = qty + value

		if value >= 90 {
			goodScore = append(goodScore, value)
		}
	}
	average := float64(qty) / float64(len(scores))
	fmt.Println(average)
	fmt.Println(goodScore)

}
