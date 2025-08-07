package main

import "fmt"

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

}
