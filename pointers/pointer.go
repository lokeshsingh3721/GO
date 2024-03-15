package main

import "fmt"

func main(){

	var a int = 10

b:= &a // refrence the value 

fmt.Println(*b) // derefrence 

*b = 30

fmt.Println(a)	// a = 30 now 

}