package main

import (
	"log"
	"net/http"
)

func main() {
	// initialize a new servemux
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/note/view", home)
	mux.HandleFunc("/note/create", noteCreate) // POST

	// web server
	log.Println("Starting server on :8000")
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
