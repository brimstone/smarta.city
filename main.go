package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":8080"
	}

	log.Println("Starting")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	})

	http.ListenAndServe(port, nil)
}
