package main

import "fmt"

func test[T any](a T){
	fmt.Println(a)
}

func main(){

	test[int](20)
	test[string]("working")


}