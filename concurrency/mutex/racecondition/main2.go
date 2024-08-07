package racecondition

import (
	"fmt"
	"sync"
	"time"
)

var GlobalSlice []int

func ProcessArray(wg *sync.WaitGroup, mut *sync.Mutex, data int) {
	mut.Lock()
	for i := 1; i <= 10; i++ {
		GlobalSlice = append(GlobalSlice, data)
	}
	mut.Unlock()
	wg.Done()
}

func PrintData() {
	for data := range GlobalSlice {
		fmt.Print(data, " ")
	}
	fmt.Println("\n\n Total size: ", len(GlobalSlice)) // there are 2 loops one run 5 and other run 10 so total ther should be 50 elements
	// with out mutex we will get length like 16, 26, 27 etc..
	time.Sleep(2 * time.Second)
}

// Output without mutex-lock

// Race Condition on Array.........
// 0 1 2 3 4 5 6 7 8 9

//  Total size:  10
// Main Terminated

// Race Condition on Array.........
// 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17

//  Total size:  18
// Main Terminated

// output with mutex lock

// Race Condition on Array.........
// 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49

//  Total size:  50
// Main Terminated
