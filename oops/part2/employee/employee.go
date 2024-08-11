package employee

import "fmt"

type Income interface {
	Source()
	Calculate() int
}

type Address struct {
	City    string
	Address string
}

func (a *Address) GetEmpAddress() {
	fmt.Printf("Emp addressdetails: %+v\n\n", a)
}

type Employee struct {
	EmpID   string
	Address // --> this is the anonaymous field, so the Employee will treat all the properties of Address as the fileds of outer struct
}

type Permanent struct {
	Employee // embedded a type to another type
	Base     int
	PF       int
	EmpType  string
}

func (emp *Permanent) Source() {
	fmt.Println("Source of income is: ", emp.EmpType)
}
func (emp *Permanent) Calculate() int {
	return emp.Base + emp.PF
}

type Contract struct {
	Employee // embedded a type to another type
	Base     int
	EmpType  string
}

func (emp *Contract) Source() {
	fmt.Println("Source of income is: ", emp.EmpType)
}

func (emp *Contract) Calculate() int {
	return emp.Base
}

func NetIncomes(incomes *[]Income) {
	total := 0
	for _, val := range *incomes {
		fmt.Print("The source of Income :---> ")
		val.Source()
		fmt.Printf("And complete details of the current emp: %+v\n", val)
		total += val.Calculate()
	}
	fmt.Println("Net company income : ", total)
}

type Freelencer struct {
	Employee
	PerHrCost int
	TotalHr   int
	EmpType   string
}

func (emp *Freelencer) Source() {
	fmt.Println("Source of income is: ", emp.EmpType)
}

func (emp *Freelencer) Calculate() int {
	return emp.PerHrCost * emp.TotalHr
}
