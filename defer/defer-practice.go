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