package main

import "fmt"

func main(){
	
	// For loop

	// for i:=0; i<10; i++ {
	// 	fmt.Println(i)
	// }

	// while loop
	
	// i:=0
	// for i<10 {
	// 	fmt.Println(i)
	// 	i++
	// }

		// shorthand to iterate over slice
		var arr = []int{1,2,3}

		for i , el := range arr{
			fmt.Printf("index is %d and el is %d \n",i, el)
		}

}