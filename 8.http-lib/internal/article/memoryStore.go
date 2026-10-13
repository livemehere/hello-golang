package article

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

	return Article{}, nil
}

func (s *ArticleMemoryStore) Save(article Article) error {
	s.articles = append(s.articles, article)
	return nil
}

func NewArticleMemoryStore() *ArticleMemoryStore {
	return &ArticleMemoryStore{}
}
