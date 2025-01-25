package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page")
	log.Printf("home route")
}

func main() {
	http.HandleFunc("/", homeRoute)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
