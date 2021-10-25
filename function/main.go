package main

import "fmt"

func main() {

	printParameter("Saya ")

	sentence := printReturn("Aku")
	fmt.Println(sentence)

	result := add(2, 3)
	fmt.Println(result)

	luas, keliling := calculate(10, 2)
	fmt.Println(luas, keliling)

	luas2, keliling2 := calculate2(10, 2)
	fmt.Println(luas2, keliling2)

}

// function basic
func printParameter(sentence string) {
	fmt.Println(sentence)
}

// function return
func printReturn(sentence string) string {
	newSentence := sentence + " Belajar Golang"
	return newSentence
}

//function opsional
func add(number1, number2 int) int {
	return number1 + number2
}

// function with 2 return
func calculate(panjang, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	return luas, keliling
}

//function with 2 return opsional
func calculate2(panjang, lebar int) (luas int, keliling int) {
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)
	return
}
