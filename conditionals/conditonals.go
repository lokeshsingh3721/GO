package main

import "fmt"

func main(){
	var age int = 18;

	if age>=18 {
		fmt.Printf("your age is %d and you can vote ", age)
	}else {
		fmt.Printf("your age is %d and you cannot vote ", age)
	}
	
}