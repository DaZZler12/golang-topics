package methods

import (
	"fmt"
)

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

func ProcessNonStruct() {
	var s myString = myString("aryan kumar")
	vowels := s.GetVowels()
	fmt.Println("All vowels are: ", *vowels)
	fmt.Print("[")
	for i, r := range *vowels {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%c", r)
	}
	fmt.Println("]")
}
