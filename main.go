package main

import (
	"fmt"
	"net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func infoHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "info page")
}

func main() {
	fmt.Println("Listening on port :3000")

	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/info", infoHandler)

	http.ListenAndServe(":3000", nil)

}
