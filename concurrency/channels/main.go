// channels in go-lang

package main

import (
	"fmt"
	"time"
)

func Process(ch chan int) {
	fmt.Println("Hello from Process routine")
	ch <- 1
	return
}

func main() {
	// chan keyword is used

	var myChan chan int // --> it's a zero valued channel zero valued i.e. nil
	fmt.Println(myChan) // will print nil

	// proper way of making channles
	myChannel2 := make(chan int)
	fmt.Println(myChannel2)
	fmt.Println("Capacity of the channelL: ", cap(myChannel2))
	fmt.Println("current length of the channel: ", len(myChannel2))

	// send and receive to a channel is blocked by-default
	// we a go-routine is sending data to channel then must be other routine who receives from it or vice-versa

	// the channel is used to block the main routine from completion,. it waits until all the other go-routines complets
	// using <- data is read and write to a channel
	// arrow pointing to a channel is writing to it, and arrow coming out of it is data reading

	// <-myChannel2  --------> data is reading from mychannel2
	// myChannel2 <- 1 ----> writing data to mychannel2

	// the below implemntation will lead to deadlock , as I am using a unbuffered chanel thus the writing to
	// myChannel2 will be blocked until some go-rotuine redas form it, as both ops are wirtten in
	// the same funciton , thus it will lead to deadlock can be avoided using buffered channel
	// or use of immdeiately invoked funciton also solves the issue

	// ------------------ CODE causing DEADLOCK -----------------------
	// myChannel2 <- 1
	// fmt.Println("New Capacity of the channelL: ", cap(myChannel2))
	// fmt.Println("New current length of the channel: ", len(myChannel2))

	// //reading form mychannle2, if we didn;t write this then it will cause deadlock
	// dataFormChan := <-myChannel2
	// fmt.Println("Reading data from channel: ", dataFormChan)

	// _____________________**********************__________________________

	go Process(myChannel2)
	fmt.Println("Reading data from channel: ", <-myChannel2)

	// we can also do like this
	go func() {
		myChannel2 <- 20
		fmt.Println("print from IIFS")
	}()
	fmt.Println("Reading data from immediately invoked funcitons: ", <-myChannel2)

	// print sum of squre and cube of digits of a  number
	time.Sleep(1 * time.Second)
	done := make(chan bool)
	fmt.Println("Calculating....")
	go CalCulate(done, 123)
	<-done

	fmt.Println("\n\nMain routine terminated")
}

func CalCulate(done chan bool, number int) {
	calculation1 := make(chan bool)
	calculation2 := make(chan bool)

	go SquareSum(number, calculation1)
	go CubeSum(number, calculation2)
	done <- (<-calculation1 && <-calculation2)
}

func SquareSum(number int, done chan bool) {
	digitChannel := make(chan int)
	go getDigits(number, digitChannel)
	sum := 0
	for digit := range digitChannel { // for loop will break when the channle is closed
		sum += (digit * digit)
	}
	fmt.Printf("sum of square of digits of num %d is: %d\n", number, sum)
	done <- true
	close(done)
}
func CubeSum(number int, done chan bool) {
	digitChannel := make(chan int)
	go getDigits(number, digitChannel)
	sum := 0
	for digit := range digitChannel {
		sum += (digit * digit * digit)
	}
	fmt.Printf("sum of cube of digits of num %d is: %d\n", number, sum)
	done <- true
	close(done)
}

func getDigits(number int, ch chan<- int) { // getDigit can only send data thus making the Bi-Directioanl into Uni-Directional
	for number > 0 {
		time.Sleep(1 * time.Second)
		fmt.Println("current Digit: ", number%10)
		ch <- number % 10
		number /= 10
	}
	close(ch) // this is most important as we are running loop on this channel, thus close
	// will inform when to break the loop,
	// otherwise the read will get blocked always leading to deadlock..
}
