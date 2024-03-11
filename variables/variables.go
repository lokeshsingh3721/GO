package main

import "fmt"

func main(){
	var first int = 30;
	var second int = 40;

	fmt.Println(first,second)

	// another way of declaring the values 

	third:= 50;
	fourth:=60;

	fmt.Println(third,fourth)

	// declare more than one variable in single line 

	fifth , sixth := 70 , 80;

	fmt.Println(fifth,sixth)

	// typecasting in go 

	var number1 float64 = 30.8;

	var number2 int = int(number1)

	fmt.Println(number2)

	/* default types which will come handy 
	bool
	string
	int
	uint32
	byte 
	rune 
	float64
	complex128
*/ 

const name string =  "Lokesh" ;
const age int = 18;

fmt.Printf("my name is %s and i am  %v years old ",name, age)

}