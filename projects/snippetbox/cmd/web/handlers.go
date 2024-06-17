package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/"{
		  http.NotFound(w,r)
			return
	}
	w.Write([]byte("Home page"))
}

func snippetView (w http.ResponseWriter, r *http.Request){
	id , err := strconv.Atoi(r.URL.Query().Get("id"))
	if err!=nil{
		http.NotFound(w,r)
		return
	}
	_,_=fmt.Fprintf(w,"snippet with the id %d",id)
}

func snippetCreate(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost{
		w.Header().Set("Allow","POST")
		http.Error(w,"invalid method",http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("snippet created successfully"))
}