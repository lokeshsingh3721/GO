package main

import "fmt"

func sum(num ...int)int{ // creating the slice of 1,2,3 --> [1,2,3]
	res:=0;
	for i:=0;i<len(num);i++{
		res+=num[i]
	}
return res
}

func main(){

	var arr = []int{1,2,3,4,5} // slice 

	fmt.Println(sum(arr...)) // spread operator --> must pass the slice 

}