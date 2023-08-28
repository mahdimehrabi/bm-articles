package article

import (
	"bm/src/apps/grpc/proto/article"
	"bm/src/domain/article/services"
	"bm/src/entities"
	"context"
	"fmt"
)

type ArticleServer struct {
	articleService *services.ArticleService
}

func (as ArticleServer) GetArticle(req *article.GetByIDReq, server article.Article_GetArticleServer) error {
	//TODO implement me
	panic("implement me")
}

func (as ArticleServer) IncreaseBuyCount(stream article.Article_IncreaseBuyCountServer) error {
	for {
		idReq, err := stream.Recv()
		if err != nil {
			stream.SendAndClose(&article.Empty{})
			return err
		}
		fmt.Println(idReq.ID)
	}
	return nil
}

func (as ArticleServer) mustEmbedUnimplementedArticleServer() {
	//TODO implement me
	panic("implement me")
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
