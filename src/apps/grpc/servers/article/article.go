package article

import (
	"bm/src/apps/grpc/proto/article"
	"bm/src/domain/article/services"
	"bm/src/entities"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
)

type ArticleServer struct {
	articleService *services.ArticleService
}

func (as ArticleServer) GetArticle(req *article.GetByIDReq, server article.Article_GetArticleServer) error {
	//TODO implement me
	panic("implement me")
}

func (as ArticleServer) IncreaseBuyCount(reqStream article.Article_IncreaseBuyCountServer) error {
	for {
		idReq, err := reqStream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return status.Errorf(codes.Internal, "Error receiving request %v", err)
		}
		articleEnt := entities.Article{
			ID: int64(idReq.ID),
		}
		count, err := as.articleService.IncreaseCount(&articleEnt)
		if err != nil {
			return status.Errorf(codes.Internal, "Error increasing count %v", err)
		}
		err = reqStream.Send(&article.BuyCount{
			ID:    idReq.ID,
			Count: count,
		})
		if err != nil {
			return status.Errorf(codes.Internal, "Error increasing count %v", err)
		}
	}
	return nil
}

func NewArticleServer(articleService *services.ArticleService) *ArticleServer {
	return &ArticleServer{articleService: articleService}
}

func (as ArticleServer) GetArticles(context context.Context, empty *article.Empty) (*article.ArticleListResponse, error) {
	articles, err := as.articleService.GetAllArticles()
	if err != nil {
		return nil, err
	}
	articleListResponse := articlesToArticleResponse(articles)

	return &articleListResponse, nil
}

func articlesToArticleResponse(articles []*entities.Article) article.ArticleListResponse {
	articleListResponse := article.ArticleListResponse{Articles: make([]*article.ArticleResponse, 0)}
	for _, a := range articles {
		articleListResponse.Articles = append(articleListResponse.Articles, &article.ArticleResponse{
			Title:    a.Title,
			ID:       a.ID,
			Body:     a.Body,
			BuyCount: a.BuyCount,
			Price:    a.Price,
			Tags:     a.Tags,
		})
	}
	return articleListResponse
}
