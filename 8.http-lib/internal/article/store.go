package article

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ArticleStore interface {
	List() ([]Article, error)
	Find(id int) (Article, error)
	Save(article Article) (Article, error)
}
