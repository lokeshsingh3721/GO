package main

import ( 
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "gitub.com/jinzhu/gorm/dialects/mysql"
	"github.com/lokeshsingh3721/bookstore/pkg/routes"

)

func main(){
	r:= mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:9010",r))
}