// waitGroup
package main

import (
	"fmt"
	"sync"
	"time"
)

func Process(wg *sync.WaitGroup, itr int) {
	fmt.Println("Hello from process routine, going to sleep & it number: ", itr)
	time.Sleep(1 * time.Second)
	fmt.Println("Weak up, going to end & itr number: ", itr)
	wg.Done()
}

func main() {

	// wait Groups are used to block the termination of the main go routine
	// do it actually block the main go-routine untill all the child routines are completed
	wg := &sync.WaitGroup{}
	counter := 3
	for counter > 0 {
		wg.Add(1)
		go Process(wg, counter)
		counter--
	}
	wg.Wait() // this will be blocked until the counter becomes 0, i.e. All the go-routines ends
	fmt.Println("terminating main routine")

}
