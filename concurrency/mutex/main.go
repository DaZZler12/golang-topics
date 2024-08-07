package main

import (
	"fmt"
	"mutex/racecondition"
	"sync"
)

// golbal variable a Shared Resource
var x = 0
var y = 0

// proper way of handling race conditions using Mutex
func Increment(wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	y += 1
	mu.Unlock()
	wg.Done()
}

func PreventRaceCondition() {
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		Increment(wg, mut)
	}
	wg.Wait()
	fmt.Println("New Updated value is: ", y)
}

func main() {
	fmt.Println("Demo for race Condition.....")
	MakeRaceCondition()

	fmt.Println("Demo for prevent Race COndition")
	PreventRaceCondition()

	//Array ops..
	fmt.Println("Race Condition on Array.........")
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go racecondition.ProcessArray(wg, mut, i)
	}
	wg.Wait()
	racecondition.PrintData()
	fmt.Println("Main Terminated")
}

func IncrementForRaceCondition(wg *sync.WaitGroup) {
	x += 1
	wg.Done()
}

func MakeRaceCondition() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go IncrementForRaceCondition(wg)
	}
	wg.Wait()
	fmt.Println("The new Updated value of Shared Resource X--> : ", x)
}

// Critical section: The section of a code that modifies shared resources is called as critical section..

// in concurrent system, multiple threads or go-routines should not access the shared resources at a time
// otherwise it will result into Race-Condition ==> It's a condition where the the outcomes depends on the order
// of execution of threads or go-routines.. i.e. how context switch happen

// In order to Avoid this, the critical section of a code is written inside Mutex.. Mututal Exclusion Locks
// So that only one thread/ Go-Routine can access the shared resource.
// So if a go-routine holds a lock then if another comes to hold the lock, then it will be blocked the previous one completes
// and the lock is freed
// so it's prseent in sync package in Golang..
// it has 2 methods Lock() and Unlock()

// Lock()

// other pieces of CODE 		----------------> this section will be only accessible by one go-routine

// Unlock()

// sample output of race Condition:

// % go run main.go
// The new Updated value of Shared Resource X--> :  20
// Main Terminated
// abhirup@192.168.1.7 /Users/abhirup/Desktop/Codes/go_codes/golang-topics/concurrency/mutex [main]
// % go run main.go
// The new Updated value of Shared Resource X--> :  19
// Main Terminated
