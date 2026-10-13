package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

var addr = fmt.Sprintf(":%d", port)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello")
	})

	mux.HandleFunc("GET /articles", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "GET articles")
	})

	mux.HandleFunc("GET /articles/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "GET article %d", id)
	})

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	log.Println("server is running on :8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
