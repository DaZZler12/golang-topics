package main

import (
	"fmt"
	"part2/employee"
)

func main() {
	pemp1 := &employee.Permanent{
		Employee: employee.Employee{
			EmpID: "PE123",
			Address: employee.Address{
				City:    "mmedkinsd",
				Address: "qwndkmnkedols",
			},
		},
		Base:    1234444,
		PF:      23434,
		EmpType: "permanent",
	}

	pemp2 := &employee.Permanent{
		Employee: employee.Employee{
			EmpID: "PE456",
			Address: employee.Address{
				City:    "koojaiwd",
				Address: "pihasdjk",
			},
		},
		Base:    5634444,
		PF:      22434,
		EmpType: "permanent",
	}

	cemp1 := &employee.Contract{
		Employee: employee.Employee{
			EmpID: "CE123",
			Address: employee.Address{
				City:    "oukswklw",
				Address: "lkatiek",
			},
		},
		Base:    7834444,
		EmpType: "contract",
	}
	fmt.Println("Address of cemp1 is: ")
	cemp1.GetEmpAddress()

	// freelencer
	femp1 := &employee.Freelencer{
		Employee: employee.Employee{
			EmpID: "CE123",
			Address: employee.Address{
				City:    "oukswklw",
				Address: "lkatiek",
			},
		},
		PerHrCost: 5000,
		TotalHr:   90,
		EmpType:   "freelencer",
	}

	var allIncomes []employee.Income
	allIncomes = append(allIncomes, pemp1, pemp2, cemp1, femp1)
	fmt.Println("All income sources address:", allIncomes)
	fmt.Println()

	employee.NetIncomes(&allIncomes)

}

// 14760000
