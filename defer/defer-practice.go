package main

import "fmt"


func first()  {

		 fmt.Printf("this is the first function ")

}

func second()string {

	return fmt.Sprintf("this is the second  function ")
	
}

func third()string {

	return fmt.Sprintf("this is the third  function ")
	
}

func main(){

 defer first()

	fmt.Println(second())
	fmt.Println(third())

}

// Result is 

// this is the second  function 
// this is the third  function 
// this is the first function


// ? Note 

// defer first execute the sorrounding function and then execute


/* 
what is first class and higher order functions ?

first class functions are the functions which can be treated as like the other values .

example:- var ans int = sum(a int , b int )

whereas , higher order functions are the functions which takes another function as an argument.

example:-

func child (x int , y int ) int {} func parent (int a , child(x int, y int ) int ) int {}
*/