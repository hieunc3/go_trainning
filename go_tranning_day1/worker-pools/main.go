package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id       int
	randomno int
}

type Result struct {
	job         Job
	sumofdigits int
}

//create buffered channels
var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	/*
		After creating the needed worker Goroutines, it waits for all the Goroutines to finish their execution by calling wg.Wait()
		After all Goroutines finish executing, it close the results channel since all Goroutines have finished their execution
		and no one else will further be writing to the results channel
	*/
	wg.Wait()
	close(results)
}

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	/*
		It will close the jobs channel after writing all jobs
	*/
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d, sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

func main() {

	noOfJobs := 100
	//create jobs on jobs Goroutine
	go allocate(noOfJobs)
	done := make(chan bool)
	//create result Goroutine -> print the result to the screen
	go result(done)
	noOfWorkers := 20
	createWorkerPool(noOfWorkers)
	//the main waits on the done channel for all results to be printed
	<-done

}

/*
	- Worker Goroutines listen for new tasks on job buffered channel
		- func worker() which reads from jobs buffered channel -> creates a Result struct using the current job -> return value of digits func -> writes the result to the result buffered
		- this func takes a WaitGroup wg as a parameter on which it will call the Done() method when all jobs have been completed
	- Once a task is complete, the result is written to the results buffered channel
*/
