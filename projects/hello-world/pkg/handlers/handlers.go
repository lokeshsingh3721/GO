package handlers

import (
	"net/http"

	"github.com/lokeshsingh3721/go/hello-world/pkg/render"
)

func Home(w http.ResponseWriter , r *http.Request){
	render.RenderTemplate(w,"home.page.templ")
}
func About(w http.ResponseWriter , r *http.Request){
		render.RenderTemplate(w,"about.page.templ")
}



