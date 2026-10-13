package article

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type CreateArticleRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ArticleHandler struct {
	store ArticleStore
}

var idSeq = 3

func NewArticleHandler(store ArticleStore) *ArticleHandler {
	return &ArticleHandler{
		store: store,
	}
}

func (h *ArticleHandler) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	list, err := h.store.List()
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(list); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *ArticleHandler) Find(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid article id", http.StatusBadRequest)
	}

	found, err := h.store.Find(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, fmt.Sprintf("not found article %d", id), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(found); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func (h *ArticleHandler) Create(w http.ResponseWriter, r *http.Request) {
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

	if err := h.store.Save(article); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(article); err != nil {
		return
	}
}
