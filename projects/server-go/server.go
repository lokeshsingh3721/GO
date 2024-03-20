package main

import (
	"encoding/json"
	"log"
	"net/http"
)


type Todo struct {
	ID   int
	Task string
}

var todos []Todo

func main() {

	http.HandleFunc("/",func(rw http.ResponseWriter , r *http.Request){

		switch r.Method {
		case "GET":
			json.NewEncoder(rw).Encode(todos)
		case "POST":
			// create variable to save todo
			var todo Todo
			// decode `r.Body` `r *http.Request` to get data from request
			// decode the result to todo
			json.NewDecoder(r.Body).Decode(&todo)
			// append new todo in todos
			todos = append(todos, todo)
			// finally, encode again create the resutl with json format
			json.NewEncoder(rw).Encode(todo)
		}
	})

	log.Fatal(http.ListenAndServe(":8080",nil))
	
}
