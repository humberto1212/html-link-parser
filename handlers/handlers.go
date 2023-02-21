package handlers

import (
	"htmlLinkParser/utils"
	"net/http"
	"text/template"
)

func ListHandler(w http.ResponseWriter, r *http.Request) {

	list := utils.LinkParser()
	tmpl := template.Must(template.ParseFiles("ADD PATH"))

	tmpl.Execute(w, list)

}
