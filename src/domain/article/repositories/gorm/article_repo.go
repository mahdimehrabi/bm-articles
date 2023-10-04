package gorm

import (
	"bm/src/domain/article/repositories"
	"bm/src/entities"
	"bm/src/infrastracture"
	"gorm.io/gorm"
)

// GormArticleRepository is a GORM-based implementation of the ArticleRepository interface.
type GormArticleRepository struct {
	db *gorm.DB
}

func (repo *GormArticleRepository) IncreaseBuyCount(article *entities.Article) (count int64, err error) {
	tx := repo.db.Begin()
	err = tx.Exec("UPDATE articles SET buy_count=buy_count+1 WHERE id=$1", article.ID).Error
	if err != nil {
		return
	}
	err = tx.Commit().Error
	if err != nil {
		return
	}
	article, err = repo.GetByID(article.ID)
	if err != nil {
		return
	}
	count = article.BuyCount
	return
}

// NewGormArticleRepository creates a new instance of GormArticleRepository.
func NewGormArticleRepository() (*GormArticleRepository, error) {
	// Create a new GormArticleRepository instance and return it
	repo := &GormArticleRepository{db: infrastracture.DB}
	return repo, nil
}

// Create adds a new article to the database.
func (repo *GormArticleRepository) Create(article *entities.Article) error {
	return repo.db.Create(article).Error
}

// GetByID retrieves an article from the database by its ID.
func (repo *GormArticleRepository) GetByID(id int64) (*entities.Article, error) {
	article := &entities.Article{}
	if err := repo.db.First(article, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, repositories.ErrArticleNotFound
		}
		return nil, err
	}
	return article, nil
}

// GetAll retrieves all articles from the database.
func (repo *GormArticleRepository) GetAll() ([]*entities.Article, error) {
	var articles []*entities.Article
	if err := repo.db.Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

// Update modifies an existing article in the database.
func (repo *GormArticleRepository) Update(article *entities.Article) error {
	return repo.db.Save(article).Error
}

// Delete removes an article from the database by its ID.
func (repo *GormArticleRepository) Delete(id int64) error {
	return repo.db.Delete(&entities.Article{}, id).Error
}
