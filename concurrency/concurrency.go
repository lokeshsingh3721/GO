package main

import "fmt"

 func child(){
	fmt.Printf("this is first child")
}

 func child2(){
		fmt.Printf("this is second child")
}

func parent(){
	fmt.Printf("this is parent")
}

func main(){

	go parent()

	child()
	child2()
	
}

// output is 

// this is a first child
// this is a second child
// this is a parent

/* Concurrency in Action:

    Goroutine: The go keyword creates a goroutine, enabling multiple pieces of code to execute concurrently.
    Non-deterministic Output: The exact order of printed messages can vary due to scheduling decisions made by the Go runtime.
    Unpredictable Execution: Despite parent() being launched first, it might not always print first. This demonstrates the flexibility and unpredictable nature of concurrent execution.
*/