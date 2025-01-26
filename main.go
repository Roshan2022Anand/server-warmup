package main

import (
	"fmt"
	// "github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// to check if the server is up
func homeRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is up and running")
}

// to warm up the server
func warmUpServer() {
	backendURL := "https://snippet-stack.onrender.com"
	res, err := http.Get(backendURL)

	url := os.Getenv("WARMUP_SERVER_URL")
	fmt.Println(url)
	if err != nil {
		log.Printf("Error while making request to backend server %s\n", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error while reading response body %s\n", err)
		return
	}
	log.Printf(" %s\n", body)
}

func main() {

	godotenv.Load()

	//to call warmUpServer func every 25 minutes
	go func() {
		for {
			warmUpServer()
			time.Sleep(25 * time.Minute)
		}
	}()

	http.HandleFunc("/", homeRoute) //route to check if the server is up

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
