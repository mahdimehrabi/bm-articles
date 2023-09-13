package repositories

import (
	"bm/src/entities"
	"errors"
)

var ErrArticleNotFound = errors.New("article not found")

// ArticleRepository defines the interface for interacting with the Article entity.
type ArticleRepository interface {
	Create(article *entities.Article) error
	GetByID(id int64) (*entities.Article, error)
	GetAll() ([]*entities.Article, error)
	Update(article *entities.Article) error
	Delete(id int64) error
	IncreaseBuyCount(article *entities.Article) (int64, error)
}
