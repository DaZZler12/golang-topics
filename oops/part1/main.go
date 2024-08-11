package main

import (
	"fmt"
	"part1/employee"
)

func main() {
	emp := employee.NewEmployee("EUw23434", "Raj", 30, 5)
	// none of the properties of emp is accessible here, we need to create getter and setter methods to access or modify them
	fmt.Printf("Newly created emp object: %+v\n", emp)

	emp.CalculateLeaves()
}
