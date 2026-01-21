package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Channels are the communication mechanism that allows goroutines to synchronize and exchange data.
A channel is a typed conduit through which you can send and receive values with the channel operator <-.
*/

/*
To create a channel, you can use the built-in make() function, specifying the type of data it will transmit.
*/

func basic_channel() {
	ch := make(chan int)
	go func() {
		ch <- 42
	}()

	value := <-ch
	fmt.Println(value)
}

/*
Channels can also be used to synchronize the execution of goroutines.
By default, channel operations block until both the sender and receiver are ready.
This behavior allows us to control the flow of execution and ensure that
certain operations are completed before others.
For example, consider a scenario where we need to calculate the sum of
two numbers concurrently. We can use two channels to send
the numbers to be added and receive the result:
*/

func addNumbers(a, b int, result chan int) {
	result <- a + b
}

func addnumb_channel() {
	ch := make(chan int)
	go addNumbers(5, 7, ch)
	sum := <-ch
	fmt.Println(sum)
}

/*
Remember to always close channels when you're done sending data to them.
Closing a channel indicates that no more values will be sent and allows
the receiver to detect when all values have been received.
*/

func close_channels() {
	ch := make(chan int)
	go func() {
		defer close(ch)
		ch <- 42
	}()

	value, ok := <-ch
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("Channel closed")
	}
}

/*
In this example, we use the close() function to close the channel after
sending the value 42. The receiver uses the second return value of the
receive operation to check if the channel is closed.
If the channel is closed, it prints "Channel closed" to the console.
*/

// goruptins and other

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}

func sum(numbers []int, result_channel chan int) {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	result_channel <- sum
	// send data of sum to result channel
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("Worker", id, "started job", job)
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		fmt.Println("Worker", id, "finished job", job)
		results <- job * 2
	}
}

func main() {
	// GOROUTINES
	go printNumbers()
	time.Sleep(3 * time.Second)

	fmt.Println("end")

	// CHANNELS
	numbers := []int{1, 2, 3, 4, 5}
	// make channel "result"
	result := make(chan int)
	go sum(numbers, result)
	// send data of "result" to "total" channel
	fmt.Println("wait for result_channel....")
	total := <-result
	fmt.Println("got a result_channel")
	fmt.Println("Sum:", total)

	/*
	   When you run this program, you will see the sum being calculated concurrently.
	   Channels provide a synchronization mechanism, ensuring that the result is only read
	   from the channel when it is available.
	*/

	// Concurrency Patterns

	jobs := make(chan int, 5)
	results := make(chan int, 5)
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			worker(id, jobs, results)
			wg.Done()
		}(i)
	}

	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println("Result:", result)
	}

}
