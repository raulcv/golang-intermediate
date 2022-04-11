package main

import "fmt"

// Worker Concurrency with channels and goroutines - Worker pools
func main() {
	tasks := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 40}
	nWorkers := 3
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))
	for i := 0; i < nWorkers; i++ {
		go Worker(i, jobs, results)
	}
	for _, task := range tasks {
		jobs <- task
	}
	close(jobs)
	for i := 0; i < len(tasks); i++ {
		<-results
	}
}
func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker with id %d started fib with %d jobs \n", id, job)
		fib := Fibonacci(job)
		fmt.Printf("Worker with id %d, job %d and fib %d \n", id, job, fib)
		results <- fib
	}
}
func Fibonacci(n int) int {
	for n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
