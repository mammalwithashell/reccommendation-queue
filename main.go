package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting http server")
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", router)
}
