package methods

import (
	"fmt"
)

type Address struct {
	city string
	pin  int
}

func (a *Address) GetAddressInfo() {
	fmt.Printf("Address Info %+v\n\n", a)
}

type Employees struct {
	empID   string
	name    string
	salary  int
	Address // promoted field
}

// A method with value as receiver type
// in this case when the method is called
// a new copy of Employees struct will be created and that will be passed to this method
// thus a method with value as recevier won't change the object of @ caller.
func (e Employees) GetDetails() {
	fmt.Printf("Employees details: %+v\n\n", e)

}

func (e *Employees) ChangeName(newName string) {
	e.name = newName
}

func (e *Employees) GetSalaryAsPointerReceiver() {
	fmt.Println("Pointer as Receiver----->  Employee name: ", e.name, " and salary:  ", e.salary)
	e.salary += 9
}

func (e Employees) GetSalaryAsValueReceiver() {
	fmt.Println("Value as Recevier----->  Employee name: ", e.name, " and salary:  ", e.salary)
}

func Helper() {
	emp := Employees{
		empID:  "123234",
		name:   "Aryan",
		salary: 1000000,
		Address: Address{
			city: "Kolkata",
			pin:  12345,
		},
	}
	emp.GetDetails()
	emp.ChangeName("Aryan Kumar")
	fmt.Printf("New updated details: %+v\n\n", emp)

	emp.GetAddressInfo() // Output: Address Info &{city:Kolkata pin:12345}

	e1 := Employees{
		empID:  "787878787",
		name:   "Lisa",
		salary: 9000000,
	}

	e1.GetSalaryAsValueReceiver()    // no change will be made to  the caller
	(&e1).GetSalaryAsValueReceiver() // no change will be made to the caller
	// in fact we cna do this
	e1.GetSalaryAsPointerReceiver()
	fmt.Printf("New updated details of e1: %+v and new salary is: %d \n\n", e1, e1.salary) // here we will see that the salary gets changes

	e2 := &Employees{
		empID:  "898976756",
		name:   "Alis",
		salary: 560700,
	} // e2 is a pointer to an object of type Employee
	(*e2).GetSalaryAsValueReceiver()
	e2.GetSalaryAsValueReceiver()

	e2.GetSalaryAsPointerReceiver()
	fmt.Printf("New updated details of e2: %+v and new salary is: %d \n\n", e2, e2.salary) // here we will see that the salary gets changes

}

// methods in Golang are a way to get some what oops property.
// it helps in logically group the behaviour of a type.
// so that the properites and the behaviour of a type
// can be grouped all together..

// The struct and it's methods needs to be defined in the same package..

// A method is same a funciton just it has recevier type which tells, that in which type this method belongs..
// the recevier type can be accessible inside the methods..
