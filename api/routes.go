package main

import "net/http"

func (a application) Route() {
	a.router.Handle("/", http.FileServer(http.Dir("./ui/static")))
	// Serve static pages
	a.router.PathPrefix("/").Handler(http.FileServer(http.Dir("./ui/static")))
}
