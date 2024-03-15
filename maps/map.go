package main

import (
	"errors"
	"fmt"
)

type user struct{
	name string 
	phoneNo int 
}

var records = make(map[string]user)


func fillmap (names []string,phoneNo []int) (map[string]user, error)  {

	if len(names) != len(phoneNo) {
		return nil,errors.New("size is not same")
	}

	for i:=0;i<len(names);i++{
		name := names[i]
		phoneNo := phoneNo[i]

		records[name] = user{
			name: name,
			phoneNo: phoneNo,
		}

	}

	return records , nil

}


func main(){

	var names = []string{"A","B","C"}
	var phoneNo = []int{123323,324234,523423432}

	result,error := fillmap(names,phoneNo)

	if(error !=nil){
		fmt.Println(error)
	}

	fmt.Println(result) // map[A:{A 123323} B:{B 324234} C:{C 523423432}]

}