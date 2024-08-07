package main

import "fmt"

type Employee struct {
	name string
	age  int
}
type AddressInfo struct {
	City string
	Pin  int
}
type Person struct {
	Name    string
	Salary  int
	Address AddressInfo
}
type PersonAnonaymous struct {
	string
	int
}

func main() {
	// assign value with the keys
	emp1 := Employee{
		name: "abc",
		// age:  22, // --> this is possible here, but not possible in the below way if initialization
	}
	fmt.Printf("emp1 data: %+v\n", emp1)

	// assign values with just values no keys
	emp2 := Employee{"pqr", 44}
	fmt.Printf("emp2 data: %+v\n", emp2)

	// Anonaymous struct type

	emp3 := struct {
		Name string
		Id   int
	}{
		Name: "hyee",
		Id:   1234,
	}
	fmt.Printf("print Anonaymous Emp3 struct: %+v\n\n", emp3)

	// pointers to a struct
	emp4 := &Employee{
		name: "Steve",
		age:  99,
	}
	fmt.Println("pointer to a strcut emp4: ", emp4)

	fmt.Println((*emp4).name) // this way
	// or
	fmt.Println(emp4.name) // no need for explicit de-referrence

	// anonyamoud fields
	pa := PersonAnonaymous{
		string: "uuuu",
		int:    23456566,
	}
	fmt.Println("accessing anonaymous fields of a struct: name --> ", pa.string, " & salary--> ", pa.int)

	// nested struct
	p1 := Person{
		Name:   "Ram",
		Salary: 1234455666,
		Address: AddressInfo{
			City: "Bangalore",
			Pin:  4455633,
		},
	}
	fmt.Printf("Nested strcut: %+v\n\n", p1)

	// promoted fileds..
	pp := PersonPrmoted{
		Name: "huiohh",
		Age:  77,
		AddressInfo: AddressInfo{
			City: "kolkata",
			Pin:  122449,
		},
	}
	fmt.Println("accessing the promoted fields directly: ", pp.City, " and ", pp.Pin)

	// compare 2 struct
	pp1 := Person{
		Name:   "rahul",
		Salary: 123,
	}

	pp2 := Person{
		Name:   "rahul",
		Salary: 123,
	}

	fmt.Println("Compering pp1 and pp2: ", pp1 == pp2)

}

type PersonPrmoted struct {
	Name string
	Age  int
	AddressInfo
}

// struct: group of properties.. can be named or anonaymous
