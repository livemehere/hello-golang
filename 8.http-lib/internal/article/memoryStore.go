package article

import (
	"fmt"
	"sync"
)

type ArticleMemoryStore struct {
	mu       sync.RWMutex
	articles []Article
	nextID   int
}

func (s *ArticleMemoryStore) List() ([]Article, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.articles, nil
}

func (s *ArticleMemoryStore) Find(id int) (Article, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, article := range s.articles {
		if id != article.ID {
			continue
		}

		return article, nil
	}

	return Article{}, fmt.Errorf("article %d not found", id)
}

func (s *ArticleMemoryStore) Save(article Article) (Article, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	article.ID = s.nextID
	s.nextID++
	s.articles = append(s.articles, article)
	return article, nil
}

func NewArticleMemoryStore() *ArticleMemoryStore {
	return &ArticleMemoryStore{}
}
