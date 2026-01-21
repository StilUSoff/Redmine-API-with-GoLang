package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"runtime/trace"
)

/*
Logging is the first step in monitoring and debugging an application.
It allows you to track the flow of execution, capture errors, and record important events.
Golang provides a built-in package called log that makes logging easy.
Here's an example of how to use it:
*/

func logging() {
	log.Println("This is a log message")
	log.Fatalf("This is a fatal error: %s", "something went wrong")
	log.Println("It will not print after fatal log")
}

/*
Profiling is the process of analyzing the performance of an application.
It helps identify bottlenecks and optimize the code for better efficiency.
Golang has a built-in profiling tool called pprof.
Here's an example of how to use it:
*/

func profiling() {
	f, err := os.Create("profile.prof") //create file profile.prof
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = pprof.StartCPUProfile(f)
	if err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	// Your application code here

	err = pprof.WriteHeapProfile(f)
	if err != nil {
		log.Fatal(err)
	}
}

/*
In the above code, we create a file to store the profiling data.
We then start profiling the CPU usage using StartCPUProfile and stop it using StopCPUProfile.
Finally, we write the heap profile using WriteHeapProfile.
You can analyze the generated profile using various tools like go tool pprof.
*/

/*
Tracing allows you to understand the execution flow of your application.
It helps identify performance issues and uncover bottlenecks.
Golang provides a built-in tracing tool called trace.
Here's an example of how to use it:
*/

func tracing() {
	f, err := os.Create("trace.out") //create file trace.out
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		log.Fatal(err)
	}
	defer trace.Stop()

	// Your application code here
}

/*
In the above code, we create a file to store the trace data.
We then start tracing using Start and stop it using Stop.
The trace data can be visualized using the go tool trace command.
*/

func main() {
	profiling()
	tracing()
	logging()
	fmt.Println("It will not print after fatal log too")
}
