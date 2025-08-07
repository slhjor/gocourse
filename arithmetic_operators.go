package main

import (
	"fmt"
	"math"
)

func main() {
	//Variable declaration
	var a, b int = 10, 3
	var result int
	result = a + b
	fmt.Println("Addition:", result)

	result = a - b
	fmt.Println("Subtraction:", result)

	result = a * b
	fmt.Println("Multiplication:", result)

	result = a / b
	fmt.Println("Division:", result)

	result = a % b
	fmt.Println("Remainder:", result)

	const p float64 = 22 / 7.0
	fmt.Println(p)

	//Overflow with signed integers
	var maxInt int64 = 9223372036854775807 // Maximum value for int64 type in Go
	fmt.Println(maxInt)

	maxInt = maxInt + 1
	fmt.Println(maxInt)

	//Overflow with unsigned integers
	var uMaxInt uint64 = 18446744073709551615 // Maximum value for uint64 type in Go
	fmt.Println(uMaxInt)

	uMaxInt = uMaxInt + 1
	fmt.Println(uMaxInt)

	// Underflow with floating point numbers
	var smallFloat float64 = 1.0e-323

	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println(smallFloat)

}
