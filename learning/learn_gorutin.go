package main

import (
	"fmt"
	"sync"
	"time"
)

func gorutine_basic() {
	go fmt.Println("Hello concurrent world")

	// если не подождать, то программа закончится, не успев вывести сообщение
	time.Sleep(100 * time.Millisecond)

	for i := 0; i < 5; i++ {
		go func(i int) { fmt.Println(i) }(i)
	}

	time.Sleep(100 * time.Millisecond)

	go fmt.Println("Hello concurrent world")

	time.Sleep(100 * time.Millisecond)

	for i := 0; i < 5; i++ {
		go func() { fmt.Println(i) }()
	}

	time.Sleep(100 * time.Millisecond)

	i := 10
	go fmt.Printf("1. Значение переменной i равно %d\n", i)
	i++
	go fmt.Printf("2. Значение переменной i равно %d\n", i)
	go func() {
		i++
		go fmt.Printf("3. Значение переменной i равно %d\n", i)
	}()
	i++
	go fmt.Printf("4. Значение переменной i равно %d\n", i)
	time.Sleep(1000000)

}

// gorutine with sync
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	for job := range jobs {
		// Perform some task
		result := job * 2

		// Send the result to the results channel
		results <- result
	}

	// Signal that this worker has completed its task
	wg.Done()
}

func basic_print(word string) (string, string) {
	fmt.Println(word)
	res := word + "end"
	res1 := word + "endd"
	return res, res1
}

func main() {

	//// try to run multiple times, "II" will be always in different place

	go basic_print("II")

	for i := 0; i < 50; i++ {
		fmt.Printf("s")
	}

	str1, str2 := basic_print("hey")
	fmt.Println(str1, str2)

	numJobs := 10
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	var wg sync.WaitGroup

	// Start the worker goroutines
	numWorkers := 5
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Enqueue the jobs
	for i := 0; i < numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	// Wait for all workers to finish
	wg.Wait()

	// Collect the results
	close(results)
	for result := range results {
		fmt.Println(result)
	}

}
