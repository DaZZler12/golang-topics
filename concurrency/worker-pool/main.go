// worker-pool
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id  int
	num int
}

type Result struct {
	job   Job
	value int
}

// buffered channels
var JobChannel = make(chan Job, 10)
var ResultChannel = make(chan Result, 10)

func calSumOfDigits(number int) int {
	sum := 0
	for number > 0 {
		sum += number % 10
		number /= 10
	}
	time.Sleep(1 * time.Second)
	return sum
}

func Worker(wg *sync.WaitGroup) {
	// here we will process from the Job Channel
	for job := range JobChannel {
		sumValue := calSumOfDigits(job.num)
		ResultChannel <- Result{job: job, value: sumValue} // Assiginig the result output to ResultChannel
	}
	wg.Done()
}
func CreateWorkerPool(noOfWorkers int) {
	var wg = &sync.WaitGroup{}
	// here we will create worker routines of totalNumber  = noOfWorkers
	for noOfWorkers > 0 {
		wg.Add(1)
		go Worker(wg)
		noOfWorkers--
	}
	wg.Wait()
	close(ResultChannel) // no more write ops on the output or result buffered channel
}

func CreateJobs(noOfJobs int) {
	// this will insert list of data into Job Channel or Input Buffer..

	for i := 1; i <= noOfJobs; i++ {
		JobChannel <- Job{id: i, num: rand.Intn(999)}
	}
	close(JobChannel) // no more data will be pushed to Job Channel
}
func PrintFromResult(done chan<- bool) {
	for output := range ResultChannel {
		fmt.Printf("JobInfo: id:%d and value: %d and Result: %d\n\n", output.job.id, output.job.num, output.value)
	}
	done <- true
}

func main() {
	startTime := time.Now()
	done := make(chan bool, 1)
	noOfJobs := 1000
	noOfWorkers := 400
	go CreateJobs(noOfJobs)
	go CreateWorkerPool(noOfWorkers)
	go PrintFromResult(done)

	<-done
	endTime := time.Now()
	fmt.Println("Total Time taken: ", endTime.Sub(startTime))

}

// woker pool is a collection of threads, waiting for jobs to be assigned to them,
// as soon as they complete the job allocated to them, they again made themselves
// ready for the new task or job..

// this coneccpt of wroker pull helps in the efficency by removing the over head of thread creation
// and destruction again and agian for shot lived task/jobs.. as well it imporives the concurrency..

// here we will create a worker of go-routines. and there task is to print the sum  of digits of a number

// 123 --> sum of digit : 6 means we will provide a list of number to the worker pool that will calculate the  sum for each number

// various componenets..
// 1. Inupt or a Job Buffer Channel
// 2. Output or a Output Result Buffer Channel
// 3. WorkerPool to create number of worker routnies
// 4. Worker to process from Job cahnnel and write into Output channel.

// As we increase the number of workers form 10 to 20 the total duration is reduced from 10sec to 5 sec

// when I have the noofjobs = 400 and worker 40 then total time = 25 sec
// and when the no of wroker becomes 400, then totlaTime becomes: 3 Sec
