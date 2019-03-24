package main

import (
	"fmt"
	"net/http"
	"html/template"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templ.ExecuteTemplate(w, "index", nil)
}

func infoHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "<h1>info page</h1>")
}

func main() {
	fmt.Println("Listening on port :3000")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/info", infoHandler)

	http.ListenAndServe(":3000", nil)

}
