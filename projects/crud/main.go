package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
		firstname string `json:"firstname"`
		lastname string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Contend-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func main(){
	r:=mux.NewRouter()

	movies = append(movies,Movie{ID:"1",Isbn: "4356" , Title: "first Movie", Director: &Director{firstname: "John",lastname:"Doe" }})

	

	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POsT")
	r.HandleFunc("/movies/{id}",updateMovie ).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")


	fmt.Printf("Starting the server at the port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))

}
