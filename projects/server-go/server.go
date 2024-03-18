package main

import (
	"fmt"
	"log"
	"net/http"
)


func main() {

	http.HandleFunc("/",func(w http.ResponseWriter , r *http.Request){
		fmt.Fprintf(w,"hello ")
	})

	http.HandleFunc("/hi",func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"hi")
	})

	log.Fatal(http.ListenAndServe(":8081",nil))
	
}
