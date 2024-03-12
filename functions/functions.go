package main

import "fmt"

func main(){

	var ans int = sum(10,12)

	fmt.Printf("ans is %d",ans)

}

func sum(a int , b int ) int   {
	return a + b;
}