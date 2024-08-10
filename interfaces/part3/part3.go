package part3

import (
	"fmt"
	"reflect"
)

type Human interface {
	Feature()
}
type Boy struct {
	Name   string
	Gender string
}

func (b *Boy) Feature() {
	fmt.Printf("Boy fetaure: %+v\n\n", *b)
}

type Girl struct {
	Name, Gender string
}

func (g *Girl) Feature() {
	fmt.Printf("Girl fetaure: %+v\n\n", *g)
}

func TypeChecker(data interface{}) {

	// will implement the type switch
	switch value := data.(type) {
	case Human:
		value.Feature()
	default:
		fmt.Println(reflect.TypeOf(value))
		fmt.Println("Invalid Type")
	}
}

func Part2() {

	b := &Boy{
		Name:   "Bob",
		Gender: "Male",
	}
	g := &Girl{
		Name:   "Lisa",
		Gender: "Female",
	}

	fmt.Println()
	TypeChecker(b)
	TypeChecker(g)

}
