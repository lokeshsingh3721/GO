package main

import "fmt"

type user struct {
	name string 
	age int
}

func getUser () (user,error) {

	// can do more with error handling 

	return user{name: "lokesh",age: 20,},nil
}

func main(){

	user,err:= getUser();

	// nill means everything is good
	if err!=nil{
		fmt.Printf("there is an error ")
	}else {
		fmt.Printf("my name is %s and my age is %d",user.name,user.age)
	}
}

