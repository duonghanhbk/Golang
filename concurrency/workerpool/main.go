package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id       int
	randomNo int
}

//type Result struct {
//	job         Job
//	sumOfDigits int
//}

var jobs = make(chan Job, 10)

//var results = make(chan Result, 10)

func wait() {
	time.Sleep(time.Second * time.Duration(rand.Intn(2)))
}

//func digits(number int) int {
//	sum := 0
//	no := number
//	for no != 0 {
//		digit := no % 10
//		sum += digit
//		no /= 10
//	}
//	fmt.Println("before digits ---------")
//	time.Sleep(2 * time.Second)
//	fmt.Println("----------- after digits ")
//	return sum
//}

func worker(worker int, wg *sync.WaitGroup) {
	for job := range jobs {
		//output := Result{
		//	job:         job,
		//	sumOfDigits: digits(job.randomNo),
		//}
		//results <- output
		fmt.Printf("Worker %d do job %d\n", worker, job.id)
		wait()
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	//close(results)
}

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomNo := rand.Intn(999)
		job := Job{i, randomNo}
		fmt.Println("Create job number: ", i)
		jobs <- job
	}
	close(jobs)
}

//
//func result(done chan bool) {
//	for result := range results {
//		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomNo, result.sumOfDigits)
//	}
//	done <- true
//}

func main() {
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	//done := make(chan bool)
	//go result(done)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	//<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")

}
