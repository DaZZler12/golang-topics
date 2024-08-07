// use of select with channels in Golang
package main

import (
	"fmt"
	"time"
)

func Server1(ch chan<- string) {
	time.Sleep(6 * time.Second)
	ch <- "hello from server one"
	close(ch)
}

func Server2(ch chan<- string) {
	time.Sleep(3 * time.Second)
	ch <- "hello from server two"
	close(ch)
}
func main() {
	server1 := make(chan string)
	server2 := make(chan string)
	go Server1(server1)
	go Server2(server2)
	fmt.Println("Ping.....")
	select {
	case data := <-server1:
		fmt.Println(data)
	case data := <-server2:
		fmt.Println(data)

		// default is used to prevent from deadlock
		// default:
		// 	fmt.Println("Defualt used to prevent deadlock")
	}

	// from the below select any random case will be triggered...
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	go Process1(ch1)
	go Process2(ch2)
	fmt.Println("Getting info about any random case selection........")
	time.Sleep(2 * time.Second)
	select {
	case a := <-ch1:
		fmt.Println(a)
	case b := <-ch2:
		fmt.Println(b)
	}
	fmt.Println("Main terminated")
}

func Process1(ch chan<- string) {
	ch <- "hyee from process1"
	close(ch)
}

func Process2(ch chan<- string) {
	ch <- "hyee from process2"
	close(ch)
}

// Select is used to select a an send/receive channel operation out of a number of channel ops, that present as cases.
// select statement is blocked until one of the channel send/receive case opration is ready..
// if we use default then it prevents deadlock, but it won't wait for none of the send/receive operations
