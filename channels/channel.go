package main

import (
	"fmt"
	"time"
)

func say(s string, ch chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		ch <- s // Send message to the channel
	}
}

func main() {
	ch := make(chan string) // Create a channel for communication
	go say("world", ch)     // Launch "say" as a goroutine (non-blocking)

	// Continue execution in main
	for i := 0; i < 3; i++ {
		fmt.Println("Hello")
	}

	// Receive messages from the channel (blocking)
	for i := 0; i < 5; i++ {
		msg := <-ch // Receive message from the channel
		fmt.Println(msg)
	}
}


/*  make channel

 ch:= make(chan type )

 ch <-- 10; // filling the channel
 ch2:= <-- ch // reading the channel

 note :- channel is blocking operating , if there is no go routine who is reading out the channel then the code will be stuck 

	ch := make(chan int , 100) create a channel buffer

	close(ch) // to close the channel

	dont send on a closed channel

 */

