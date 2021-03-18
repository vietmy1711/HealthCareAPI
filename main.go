package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func main() {
	handleRequests()
}
