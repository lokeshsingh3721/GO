package main

import (
	"net/http"

	"github.com/lokeshsingh3721/go/hello-world/pkg/handlers"
)

const port = ":8080"


func main(){
	
http.HandleFunc("/",handlers.Home)
http.HandleFunc("/about",handlers.About)
_= http.ListenAndServe(port,nil)
}