package services

import (
	"bm/src/domain/article/repositories"
	"bm/src/entities"
	"errors"
	"fmt"
	"go.uber.org/zap"
)

// ArticleServiceImpl is an implementation of the ArticleService interface.
type ArticleServiceImpl struct {
	articleRepo repositories.ArticleRepository
	logger      *zap.Logger
}

// NewArticleService creates a new instance of ArticleServiceImpl.
func NewArticleService(articleRepo repositories.ArticleRepository, logger *zap.Logger) *ArticleServiceImpl {
	return &ArticleServiceImpl{
		articleRepo: articleRepo,
		logger:      logger,
	}
}

// CreateArticle adds a new article.
func (s *ArticleServiceImpl) CreateArticle(article *entities.Article) error {
	if article == nil {
		return errors.New("article cannot be nil")
	}

	if err := s.articleRepo.Create(article); err != nil {
		s.logger.Error("Failed to create article", zap.Error(err))
		return fmt.Errorf("failed to create article: %w", err)
	}

	return nil
}

// GetArticleByID retrieves an article by its ID.
func (s *ArticleServiceImpl) GetArticleByID(id int) (*entities.Article, error) {
	article, err := s.articleRepo.GetByID(id)
	if err != nil {
		s.logger.Error("Failed to get article by ID", zap.Int("id", id), zap.Error(err))
		return nil, fmt.Errorf("failed to get article by ID: %w", err)
	}

	return article, nil
}

// GetAllArticles retrieves all articles.
func (s *ArticleServiceImpl) GetAllArticles() ([]*entities.Article, error) {
	articles, err := s.articleRepo.GetAll()
	if err != nil {
		s.logger.Error("Failed to get all articles", zap.Error(err))
		return nil, fmt.Errorf("failed to get all articles: %w", err)
	}

	return articles, nil
}

// UpdateArticle modifies an existing article.
func (s *ArticleServiceImpl) UpdateArticle(article *entities.Article) error {
	if article == nil {
		return errors.New("article cannot be nil")
	}

	if err := s.articleRepo.Update(article); err != nil {
		s.logger.Error("Failed to update article", zap.Error(err))
		return fmt.Errorf("failed to update article: %w", err)
	}

	return nil
}

// DeleteArticle removes an article by its ID.
func (s *ArticleServiceImpl) DeleteArticle(id int) error {
	if err := s.articleRepo.Delete(id); err != nil {
		s.logger.Error("Failed to delete article", zap.Int("id", id), zap.Error(err))
		return fmt.Errorf("failed to delete article: %w", err)
	}

	return nil
}
