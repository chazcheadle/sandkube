package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		log.Fatalf("Error getting hostname.", err)
	}
	fmt.Fprintf(w, "Hello from: %s", name)
}

func main() {
	var addr string = ":8080"
	handler := http.HandlerFunc(HelloServer)
	log.Printf("Starting server on port %s", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("Could not listen on port %s", addr, err)
	}
}