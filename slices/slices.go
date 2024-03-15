package main

import "fmt"

func main () {

	// var arr =  []int{1,2,3,4,5}
	// arr2:= arr[1:4] // [inclusive , exclusive ]

	// // arr2[0] = 100 // [reference to the original array ]
	
	// fmt.Println(arr)

	arr2:= make([]int , 3 ) // created the slice

	arr2[0] = 10
	arr2[1] = 10
	arr2[2] = 10
	arr2[3] = 10



	fmt.Println(arr2)

}

