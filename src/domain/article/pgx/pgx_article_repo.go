package pgx

import (
	"bm/src/entities"
	"context"
	"github.com/jackc/pgx/v4"
)

// PgxArticleRepository is a PGX-based implementation of the ArticleRepository interface.
type PgxArticleRepository struct {
	db *pgx.Conn
}

// NewPgxArticleRepository creates a new instance of PgxArticleRepository.
func NewPgxArticleRepository(db *pgx.Conn) *PgxArticleRepository {
	return &PgxArticleRepository{db: db}
}

// Create adds a new article to the database.
func (repo *PgxArticleRepository) Create(article *entities.Article) error {
	_, err := repo.db.Exec(context.Background(), "INSERT INTO articles (title, body, tags, price) VALUES ($1, $2, $3, $4)", article.Title, article.Body, article.Tags, article.Price)
	return err
}

// GetByID retrieves an article from the database by its ID.
func (repo *PgxArticleRepository) GetByID(id int) (*entities.Article, error) {
	row := repo.db.QueryRow(context.Background(), "SELECT title, body, tags, price FROM articles WHERE id = $1", id)
	article := &entities.Article{}
	err := row.Scan(&article.Title, &article.Body, &article.Tags, &article.Price)
	if err != nil {
		return nil, err
	}
	return article, nil
}

// GetAll retrieves all articles from the database.
func (repo *PgxArticleRepository) GetAll() ([]*entities.Article, error) {
	rows, err := repo.db.Query(context.Background(), "SELECT title, body, tags, price FROM articles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []*entities.Article
	for rows.Next() {
		article := &entities.Article{}
		err := rows.Scan(&article.Title, &article.Body, &article.Tags, &article.Price)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}

// Update modifies an existing article in the database.
func (repo *PgxArticleRepository) Update(article *entities.Article) error {
	_, err := repo.db.Exec(context.Background(), "UPDATE articles SET title=$1, body=$2, tags=$3, price=$4 WHERE id=$5", article.Title, article.Body, article.Tags, article.Price)
	return err
}

// Delete removes an article from the database by its ID.
func (repo *PgxArticleRepository) Delete(id int) error {
	_, err := repo.db.Exec(context.Background(), "DELETE FROM articles WHERE id = $1", id)
	return err
}
