package main

import (
	"log"
	"net/http"
)

func main() {
	const port = "8080"

	mu := http.NewServeMux()

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mu,
	}

	log.Printf("Serving on port:%s\n", port)
	log.Fatal(server.ListenAndServe())
}
