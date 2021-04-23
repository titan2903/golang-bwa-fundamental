package main

import (
	"fmt"
	"go_bwa/calculation"
	"go_bwa/multiply"
)

func main() {
	// fmt.Println("Titan")

	sentence := Testaja()
	fmt.Println(sentence)
	result := calculation.Add(9, 8)

	fmt.Println(result)

	new_result := multiply.Multiply(5, 6)
	fmt.Println(new_result)

	age := 20
	age = 50
	fmt.Println(age, "=== age ===")

}
