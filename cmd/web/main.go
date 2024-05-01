package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// accept and parse command-line flag
	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	// initialize a new servemux
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/note/view", noteView)
	mux.HandleFunc("/note/create", noteCreate) // POST

	// web server
	log.Printf("Starting server on %s", *addr) // INFO
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err) // ERROR
}
