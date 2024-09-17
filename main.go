package main

import (
	"fmt"
	"net/http"

	"log"
)

func main() {
	http.HandleFunc("/", HelloWorld)
	log.Printf("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello, World!")
}
