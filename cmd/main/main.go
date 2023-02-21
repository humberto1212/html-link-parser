package main

import (
	"fmt"
	"htmlLinkParser/handlers"
	"htmlLinkParser/utils"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	results := utils.LinkParser()
	fmt.Println(results)

	//================
	//	Handlers
	//================
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.ListHandler).Methods("GET")

	// Start the HTTP server
	http.ListenAndServe(":8000", r)
}
