// unbufferd_channels

package main

import (
	"fmt"
	"time"
)

func Process(ch chan<- int) { // it will a uni-directional channel
	for i := 1; i <= 5; i++ {
		ch <- i // so after adding 2 values the ch channel will be bliocked as the capacity is 2
		fmt.Println("write to channel value: ", i)
		// so until data is read form the ch it will get blocked
	}
	close(ch) // it will ackonwledge the for loop that no more data is coming and
	// will break the loop
}

func main() {
	// send to an unbufferd channle is bloicked when the channle is full
	// recieve from an unbufferd channel is blocked when the channel is empty

	ch := make(chan int, 2) // creating a buffered channel of size 2 , so atmax it can accomodate 2 integer
	fmt.Println("channel capacity: ", cap(ch))
	fmt.Println("current length: ", len(ch))
	// ----> if we run at this point we won't get deadlock as the buffered channels are not blocked by-default

	ch <- 1
	ch <- 2
	fmt.Println("data 1 from ch: ", <-ch)
	fmt.Println("data 2 from ch: ", <-ch)

	tempCh := make(chan int, 2)
	go Process(tempCh)
	// read and print the temCh
	time.Sleep(1 * time.Second)
	for data := range tempCh {
		fmt.Println("Data received: ", data)
	}

	// we can also read froan a closed buffer channel
	// until all the data inserted into it is read completely

	newch := make(chan int, 3)
	for i := 1; i < 4; i++ {
		newch <- i
	}
	close(newch)

	// still we can read teh instered data it
	for data := range newch {
		fmt.Println("data from newch channel: ", data)
	}
	data, ok := <-newch
	fmt.Printf("The current status of newch channle: %d,%v\n", data, ok)
}
