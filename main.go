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

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	})

	http.Handle("/", http.FileServer(http.Dir("static")))

	http.ListenAndServe(port, nil)
}
