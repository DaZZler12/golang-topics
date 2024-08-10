package part4

import "fmt"

type Human interface {
	GetGender()
}

type Boy struct {
	name, gender string
}

func (b Boy) GetGender() {
	fmt.Println("I am a boy gender: ", b.gender, " and name : ", b.name)
}

type Girl struct {
	name, gender string
}

func (g *Girl) GetGender() {
	fmt.Println("I am a girl gender: ", g.gender, " and name : ", g.name)
}

func Processpart4() {
	fmt.Println("\n\n Interface with Method having pointer and value as receiver types.............................")
	// here will discuss about
	// interface with value as receiver and pointer as receivere types..
	boy := Boy{name: "Bob", gender: "male"}
	girl := Girl{name: "Alic", gender: "female"}

	var h Human
	h = boy
	h.GetGender()

	// h = girl // --> this will create a issue
	// cannot use girl (variable of type Girl) as Human value in assignment:
	// Girl does not implement Human (method GetGender has pointer receiver)compilerInvalidIfaceAssign
	// this because the instace girl when assigned to the interface then it's not implemntating
	// the method as pointer as receiver..
	// since the undelying data of an interface is not addressable so the compiler is unable to get the address and call
	// the method with receiver as pointer type..

	h = &girl // -->> this is the right way to do this
	h.GetGender()

}
