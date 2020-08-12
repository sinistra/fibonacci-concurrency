package main

import (
	"fmt"
	"time"
)

func main() {
	defer elapsed()()
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// concurrent workers
	// remove or add some to test different configs
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	// go worker(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 50; j++ {
		fmt.Println(<-results)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fibonacci(n)
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// calculate time elapsed
func elapsed() func() {
	start := time.Now()
	return func() {
		fmt.Printf("Calculation took %v\n", time.Since(start))
	}
}
