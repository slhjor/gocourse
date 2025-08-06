package main

import "fmt"

type EmployeeGoogle struct {
	FirstName string
	LastName  string
	Age       int
}

type EmployeeApple struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	//  PascalCase
	//Eg. CalculateArea, UserInfo, NewHttpRequest
	//  Structs, interfaces, enums

	// snake_case
	//Eg. user_id, first_name, http_request

	// UPPERCASE
	//Eg. Use case is Constants

	// mixedCase
	//Eg. javaScript, htmlDocument, isValid

	const MAXRETRIES = 5

	var employeeID = 1001
	fmt.Println("EmployeeID: ", employeeID)

}
