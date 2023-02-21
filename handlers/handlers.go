package handlers

import (
	"net/http"
	"text/template"
)

func ListHandler(w http.ResponseWriter, r *http.Request) {

	list := "NEED TO IMPORT VALUE"
	tmpl := template.Must(template.ParseFiles(""))

	tmpl.Execute(w, list)

}
