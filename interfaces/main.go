package main

import (
	"fmt"
	"interfaces/part2"
	"interfaces/part3"
	"interfaces/part4"
)

type VowelsFinder interface {
	GetVowels() *[]rune
}

type myString string

// logic is made to handle only string with small case
func (s myString) GetVowels() *[]rune {
	var vowels []rune
	for _, c := range s {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			vowels = append(vowels, c)
		}
	}
	return &vowels
}

// now we can say that the type myString implements VowelsFinder interface

func main() {
	var s VowelsFinder = myString("aryan kumar") // here the object is of type mystring but assigned to a interface type..
	// same as oops , where we an have the object of type interface,
	// with a concrete type and value.. this can be done when a class or
	// a struct implemets an interface

	for _, val := range *s.GetVowels() {
		fmt.Printf("%c ", val)
	}
	fmt.Println()

	fmt.Println("Demonstrate Interfaces using Employeesss Data .........................")
	PEmp1 := &PEmployee{
		id:      "hhkhsef78787",
		basePay: 134455666,
		pf:      38849,
	}
	PEmp2 := &PEmployee{
		id:      "jupahiher1293",
		basePay: 892370973,
		pf:      28485,
	}
	CEmp1 := &CEmployee{
		id:      "gjyjak9986wm",
		basePay: 27748209,
	}
	CEmp2 := &CEmployee{
		id:      "mmhswik9977s",
		basePay: 9940903,
	}

	freeLencer1 := &FreeLencer{
		id:        "Feu84y89937",
		costperhr: 250,
		totalhr:   90,
	}
	freeLencer2 := &FreeLencer{
		id:        "Fek984570",
		costperhr: 250,
		totalhr:   90,
	}
	var AllEmployes []CompanySalaryCalculator = []CompanySalaryCalculator{PEmp1, PEmp2, CEmp1, CEmp2, freeLencer1, freeLencer2}
	var totalExpenditure = CalculateCompanyExpenditure(&AllEmployes)
	fmt.Println("TOtal Expenditure is: ", totalExpenditure)

	// Discussion of Empty Interfaces and Type Assertions..
	fmt.Println("\nDiscussion on empty interfaces...........................\n\n")
	part2.EmptyInterfaces()

	part3.Part2()

	part4.Processpart4()

}

// Another demonstration of Real life example of Interfaces
type CompanySalaryCalculator interface {
	CalculateSalary() int
}
type PEmployee struct {
	id      string
	basePay int
	pf      int
}

func (emp *PEmployee) CalculateSalary() int {
	return emp.basePay + emp.pf
}

type CEmployee struct {
	id      string
	basePay int
}

func (emp *CEmployee) CalculateSalary() int {
	return emp.basePay
}

type FreeLencer struct {
	id        string
	totalhr   int
	costperhr int
}

func (emp *FreeLencer) CalculateSalary() int {
	return emp.costperhr * emp.totalhr
}

// this will clauclate the total salary for all types of employees
func CalculateCompanyExpenditure(emplyes *[]CompanySalaryCalculator) int {
	total := 0

	for _, emp := range *emplyes {
		fmt.Printf("\nCalculate Salary for emp details: %+v and the type is: %T\n", emp, emp)
		total += emp.CalculateSalary() // so emp object wise differnt CalculateSalary() method will be called..
		// so in future a new type of employee is added then there won't be any cange required in the
		// complete company expenditure logic..
	}
	return total

}

// it's a collection of function signature..
// a type can implement an interface if
// it give description for all the methods inside the interface
