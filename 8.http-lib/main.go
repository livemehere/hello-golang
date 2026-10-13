package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const port = 8080

var (
	addr  = fmt.Sprintf(":%d", port)
	idSeq = 3
)

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateArticleRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articles = []Article{
	{
		ID:      1,
		Title:   "heloo",
		Content: "first articles",
	},
	{
		ID:      2,
		Title:   "golang",
		Content: "net/http is good",
	},
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /articles", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(articles); err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("POST /articles", func(w http.ResponseWriter, r *http.Request) {
		var input CreateArticleRequest
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}

		idSeq++
		article := Article{
			ID:      idSeq,
			Title:   input.Title,
			Content: input.Content,
		}

		articles = append(articles, article)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(article); err != nil {
			return
		}
	})

	mux.HandleFunc("GET /articles/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "invalid article id", http.StatusBadRequest)
		}

		for _, article := range articles {
			if article.ID != id {
				continue
			}
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(article); err != nil {
				http.Error(w, "failed to encode response", http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		http.Error(w, fmt.Sprintf("not found article %d", id), http.StatusInternalServerError)
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
