package greetings

import "fmt"

func Hello(name string) string {
message := fmt.Sprintf("hi , %v welcome ", name) // := means declaring and initializing 
return message
}