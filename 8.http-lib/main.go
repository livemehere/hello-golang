package main

import (
	"fmt"
	"log"
	"net/http"

	"http-lib/internal/article"
)

const port = 8080

var addr = fmt.Sprintf(":%d", port)

func main() {
	mux := http.NewServeMux()

	store := article.NewArticleMemoryStore()
	handler := article.NewArticleHandler(store)

	mux.HandleFunc("GET /articles", handler.List)
	mux.HandleFunc("POST /articles", handler.Create)
	mux.HandleFunc("GET /articles/{id}", handler.Find)

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	log.Println("server is running on :8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
