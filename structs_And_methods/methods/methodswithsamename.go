package methods

import "fmt"

type Rectangle struct {
	Length int
	Width  int
}

func (r *Rectangle) Area() {
	fmt.Println("Area of the Rectangle: ", r.Width*r.Length)
}

type Circle struct {
	Radius int
}

func (c *Circle) Area() {
	fmt.Println("Area of the Circle: ", 3.147*float64(c.Radius)*float64(c.Radius))
}

func ProcessShape() {
	r := &Rectangle{
		Length: 10,
		Width:  12,
	}
	r.Area()
	c := &Circle{
		Radius: 4,
	}
	c.Area()
}

// we ca have methods with same name but they should have differnet receiver..
// In funciton we can't have that..
// over loading is ot there in Go with normal funciton
// but we can have in structs
