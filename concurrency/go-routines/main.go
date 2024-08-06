package main

import (
	"fmt"
	"time"
)

func Process() {
	fmt.Println("print 1")
}

func PrintNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func PrintAlphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func main() {
	go Process()
	time.Sleep(1 * time.Second) // without this line we will get only main funciton output,
	// as the call to go-routines is retutned immediately, and all return values are ignored
	// so after the line 13 teh control immediately comes to the 14th line without waiting for the
	// Process Go-Rotuine to complete execution and come-back.. thus without the sleep, the main go-routine will
	// be completed and once the main go routine ends then all other go routines will end..

	go PrintNumbers()
	go PrintAlphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Printf("\nmain routine terminated\n")
}

// 									Go-Routines
// Process															main
// main go routine ends then the Application ends..

// the use of Sleep is not correct this need to be better handled using Go Channels,
// They are like pipe by which communication happens betwn 2 or more go routines..
