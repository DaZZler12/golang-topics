package employee

import "fmt"

type EmployeeOps interface {
	CalculateLeaves()
}

type employee struct { // now the struct will be inaccessible outisde this package, in order to access it we can use New() methdo
	id          string
	name        string
	totalLeaves int
	leaveTaken  int
}

func (emp *employee) CalculateLeaves() {
	fmt.Println("Leave balance: ")
	fmt.Println("Taken: ", emp.leaveTaken)
	fmt.Println("Left out: ", emp.totalLeaves-emp.leaveTaken)
}

func NewEmployee(id, name string, totalLeave, leavestaken int) employee {
	// here we can implement singleton pattern in order to avoid creartion of multiple instaces, if needed
	emp := employee{
		id:          id,
		name:        name,
		totalLeaves: totalLeave,
		leaveTaken:  leavestaken,
	}
	return emp
}
