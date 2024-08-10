package part2

import "fmt"

func describe(i interface{}) {
	fmt.Printf("getting details of interface: %v and the type is: %T\n\n", i, i)
}

func AssertType(data interface{}) {
	assertedData, ok := data.(int) // this will prevent the routine form panic, in case of Invalid type assertion
	if ok {
		fmt.Println("--------------->   assertion true and type is int: ", assertedData)
	}

	// type assertion using cases
	switch value := data.(type) { // by thid we will get the concreate type of the interface
	// so this data.(type) --> is only useful in case of type switch
	// so it will tell the concreate type of the interface in runtime

	case int:
		var intData int = int(value)
		fmt.Println("concreate type is int")
		fmt.Println("value after type assertion: ", intData)
	case string:
		var stringData string = string(value)
		fmt.Println("concreate type is string")
		fmt.Println("value after type assertion: ", stringData)
	case bool:
		var boolData bool = bool(value)
		fmt.Println("concreate type is string")
		fmt.Println("value after type assertion: ", boolData)
	default:
		fmt.Println("Invalid type entered")

	}
}

func EmptyInterfaces() {
	describe(12)
	describe("nmlndef iashd")
	describe([]int{1, 2, 3, 4})
	describe([]string{"jndejf", "lopjjsd", "orwyhas"})
	describe(9098.78)
	describe(true)
	describe(struct {
		name string
		age  int
	}{
		name: "Aryn Kumar",
		age:  19,
	},
	)

	AssertType(89)
	AssertType(false)
	AssertType("nknsdf")
	AssertType(99.78)
}
