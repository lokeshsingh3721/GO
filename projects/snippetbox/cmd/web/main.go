package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/",home)
	mux.HandleFunc("/snippet/view",snippetView) 
	mux.HandleFunc("/snippet/create",snippetCreate)

fmt.Println("server is listening to the port 8080")
err:=	http.ListenAndServe(":8080",mux)
if err != nil{
	log.Fatal(err)
}
}