package article

import "bm/src/entities"

// ArticleRepository defines the interface for interacting with the Article entity.
type ArticleRepository interface {
	Create(article *entities.Article) error
	GetByID(id int) (*entities.Article, error)
	GetAll() ([]*entities.Article, error)
	Update(article *entities.Article) error
	Delete(id int) error
}
