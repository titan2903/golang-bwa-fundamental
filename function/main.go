package main

import (
	"errors"
	"fmt"
)

func main() {
	// sentence := printMyResult("saya belajar")
	// fmt.Println(sentence)
	// 	result_add := add(30, 30)
	// 	fmt.Println(result_add)

	// luas, keliling := calculate(10, 2)
	// fmt.Println(luas)
	// fmt.Println(keliling)
	// scores := []int{10, 5, 8, 9, 7}
	// total := sum(scores)
	// fmt.Println(total)

	result, err := calculate(10, 2, "&")
	if err != nil {
		fmt.Println("terjadi kesalahan ")
		fmt.Println(err.Error())
	}

	fmt.Println(result)
}

func calculate(number, numberTwo int, operation string) (int, error) {
	var result int
	var errorResult error

	switch operation {
	case "+":
		result = number + numberTwo
	case "-":
		result = number - numberTwo
	case "/":
		result = number / numberTwo
	case "*":
		result = number * numberTwo
	default:
		errorResult = errors.New("Unknown operation")
	}

	return result, errorResult
}

// func sum(numbers []int) int {
// 	total := 0
// 	for _, number := range numbers {
// 		total += number
// 	}

// 	return total
// }

// func calculate(panjang int, lebar int) (luas int, keliling int) {
// 	// luas := panjang * lebar
// 	// keliling := 2 * (panjang + lebar)
// 	luas = panjang * lebar
// 	keliling = 2 * (panjang + lebar)
// 	return
// }

// func add(number int, numberTwo int) int {
// 	// result := number + numberTwo
// 	return number + numberTwo
// }

// func printMyResult(sentense string) string {

// 	newSentence := sentense + " golang"
// 	return newSentence
// }

// 1. Input
// 2. Proses //! Pasti ada
//  3> Output
