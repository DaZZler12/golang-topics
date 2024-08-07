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
	noOfJobs := 100
	noOfWorkers := 40
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

// sample output:

// JobInfo: id:28 and value: 964 and Result: 19

// JobInfo: id:18 and value: 506 and Result: 11

// JobInfo: id:24 and value: 298 and Result: 19

// JobInfo: id:30 and value: 705 and Result: 12

// JobInfo: id:1 and value: 878 and Result: 23

// JobInfo: id:40 and value: 189 and Result: 18

// JobInfo: id:22 and value: 207 and Result: 9

// JobInfo: id:3 and value: 407 and Result: 11

// JobInfo: id:36 and value: 152 and Result: 8

// JobInfo: id:33 and value: 734 and Result: 14

// JobInfo: id:34 and value: 203 and Result: 5

// JobInfo: id:35 and value: 840 and Result: 12

// JobInfo: id:38 and value: 718 and Result: 16

// JobInfo: id:37 and value: 357 and Result: 15

// JobInfo: id:14 and value: 362 and Result: 11

// JobInfo: id:12 and value: 538 and Result: 16

// JobInfo: id:13 and value: 750 and Result: 12

// JobInfo: id:16 and value: 215 and Result: 8

// JobInfo: id:15 and value: 436 and Result: 13

// JobInfo: id:29 and value: 942 and Result: 15

// JobInfo: id:21 and value: 272 and Result: 11

// JobInfo: id:19 and value: 20 and Result: 2

// JobInfo: id:20 and value: 914 and Result: 14

// JobInfo: id:26 and value: 565 and Result: 16

// JobInfo: id:17 and value: 630 and Result: 9

// JobInfo: id:27 and value: 43 and Result: 7

// JobInfo: id:39 and value: 84 and Result: 12

// JobInfo: id:23 and value: 266 and Result: 14

// JobInfo: id:32 and value: 249 and Result: 15

// JobInfo: id:2 and value: 636 and Result: 15

// JobInfo: id:25 and value: 135 and Result: 9

// JobInfo: id:4 and value: 983 and Result: 20

// JobInfo: id:31 and value: 562 and Result: 13

// JobInfo: id:5 and value: 895 and Result: 22

// JobInfo: id:6 and value: 735 and Result: 15

// JobInfo: id:7 and value: 520 and Result: 7

// JobInfo: id:8 and value: 998 and Result: 26

// JobInfo: id:11 and value: 212 and Result: 5

// JobInfo: id:9 and value: 904 and Result: 13

// JobInfo: id:10 and value: 150 and Result: 6

// JobInfo: id:51 and value: 801 and Result: 9

// JobInfo: id:42 and value: 256 and Result: 13

// JobInfo: id:43 and value: 928 and Result: 19

// JobInfo: id:47 and value: 97 and Result: 16

// JobInfo: id:48 and value: 694 and Result: 19

// JobInfo: id:49 and value: 370 and Result: 10

// JobInfo: id:41 and value: 871 and Result: 16

// JobInfo: id:50 and value: 102 and Result: 3

// JobInfo: id:75 and value: 420 and Result: 6

// JobInfo: id:72 and value: 450 and Result: 9

// JobInfo: id:80 and value: 512 and Result: 8

// JobInfo: id:76 and value: 144 and Result: 9

// JobInfo: id:77 and value: 766 and Result: 19

// JobInfo: id:78 and value: 382 and Result: 13

// JobInfo: id:79 and value: 160 and Result: 7

// JobInfo: id:73 and value: 683 and Result: 17

// JobInfo: id:74 and value: 500 and Result: 5

// JobInfo: id:71 and value: 829 and Result: 19

// JobInfo: id:44 and value: 252 and Result: 9

// JobInfo: id:45 and value: 55 and Result: 10

// JobInfo: id:46 and value: 711 and Result: 9

// JobInfo: id:52 and value: 403 and Result: 7

// JobInfo: id:53 and value: 738 and Result: 18

// JobInfo: id:54 and value: 690 and Result: 15

// JobInfo: id:55 and value: 13 and Result: 4

// JobInfo: id:56 and value: 403 and Result: 7

// JobInfo: id:57 and value: 728 and Result: 17

// JobInfo: id:58 and value: 606 and Result: 12

// JobInfo: id:59 and value: 871 and Result: 16

// JobInfo: id:60 and value: 14 and Result: 5

// JobInfo: id:61 and value: 728 and Result: 17

// JobInfo: id:62 and value: 213 and Result: 6

// JobInfo: id:63 and value: 579 and Result: 21

// JobInfo: id:64 and value: 885 and Result: 21

// JobInfo: id:65 and value: 487 and Result: 19

// JobInfo: id:66 and value: 638 and Result: 17

// JobInfo: id:67 and value: 764 and Result: 17

// JobInfo: id:68 and value: 656 and Result: 17

// JobInfo: id:69 and value: 122 and Result: 5

// JobInfo: id:70 and value: 164 and Result: 11

// JobInfo: id:87 and value: 236 and Result: 11

// JobInfo: id:100 and value: 641 and Result: 11

// JobInfo: id:91 and value: 515 and Result: 11

// JobInfo: id:92 and value: 40 and Result: 4

// JobInfo: id:93 and value: 627 and Result: 15

// JobInfo: id:94 and value: 766 and Result: 19

// JobInfo: id:95 and value: 450 and Result: 9

// JobInfo: id:96 and value: 922 and Result: 13

// JobInfo: id:97 and value: 732 and Result: 12

// JobInfo: id:98 and value: 315 and Result: 9

// JobInfo: id:99 and value: 961 and Result: 16

// JobInfo: id:90 and value: 343 and Result: 10

// JobInfo: id:81 and value: 366 and Result: 15

// JobInfo: id:86 and value: 421 and Result: 7

// JobInfo: id:88 and value: 167 and Result: 14

// JobInfo: id:89 and value: 728 and Result: 17

// JobInfo: id:83 and value: 988 and Result: 25

// JobInfo: id:82 and value: 981 and Result: 18

// JobInfo: id:84 and value: 595 and Result: 19

// JobInfo: id:85 and value: 47 and Result: 11

// Total Time taken:  3.003473958s
