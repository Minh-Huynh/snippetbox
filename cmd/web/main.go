package main

import (
    "log"
    "net/http"
	"flag"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
    mux := http.NewServeMux()
	//This fileServer serves files relative to the static directory
	//We have to strip the prefix, otherwise, it'll look for a file
	//in ui/static/static/ 
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	//mux.Handle("/static/", fileServer)
    mux.HandleFunc("/", home)
    mux.HandleFunc("/snippet/view", snippetView)
    mux.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("Starting server on %s", *addr)
    err := http.ListenAndServe(*addr, mux)
    log.Fatal(err)
}