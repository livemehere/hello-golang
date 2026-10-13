package article

import "fmt"

type ArticleMemoryStore struct {
	articles []Article
}

func (s *ArticleMemoryStore) List() ([]Article, error) {
	return s.articles, nil
}

func (s *ArticleMemoryStore) Find(id int) (Article, error) {
	for _, article := range s.articles {
		if id != article.ID {
			continue
		}

		return article, nil
	}

	return Article{}, fmt.Errorf("article %d not found", id)
}

func (s *ArticleMemoryStore) Save(article Article) error {
	s.articles = append(s.articles, article)
	return nil
}

func NewArticleMemoryStore() *ArticleMemoryStore {
	return &ArticleMemoryStore{}
}
