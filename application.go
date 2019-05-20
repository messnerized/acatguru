package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	http.Handle("/", http.FileServer(http.Dir("./")))
	log.Printf("Started webserver on port %s\n\n", port)
	http.ListenAndServe(":" + port, nil)
}

