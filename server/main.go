package main

import (
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"net/http"
)

func main(){

	fmt.Println("hjello")
	router := mux.NewRouter()

	// Return 10 random phrase pairs
	router.HandleFunc("/api/1.0/phrases", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	}).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}