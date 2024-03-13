package main

import (
	"fmt"
)


func sum(a int , b int ) int   {
	return a + b;
}

// function can take another function as an arguement

func child (b string) string {
	return b;
}

func parent (a  string ,child func (string) string) string {
	res:=child(a);
	return res
}


func main(){

	// var ans int = sum(10,12)

	// fmt.Printf("ans is %d",ans)

fmt.Printf(parent("lokesh",child))

	

}



