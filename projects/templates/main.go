package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const port = ":8080"

func Home(w http.ResponseWriter , r *http.Request){
	renderTemplate(w,"home.page.tmpl")
}
func About(w http.ResponseWriter , r *http.Request){
		renderTemplate(w,"about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, templ string){
	parsedTemplate,_ := template.ParseFiles("./pages/" + templ)
	  err:= parsedTemplate.Execute(w,nil)
if err != nil{
	_,_=fmt.Fprintf(w,fmt.Sprintf("there is some error %s",err))
}
}


func main(){
	
http.HandleFunc("/",Home)
http.HandleFunc("/about",About)
_= http.ListenAndServe(port,nil)
}