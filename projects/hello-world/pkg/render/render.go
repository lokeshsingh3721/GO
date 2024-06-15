package render

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, templ string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + templ)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		_, _ = fmt.Fprintf(w, fmt.Sprintf("there is some error %s", err))
	}
}