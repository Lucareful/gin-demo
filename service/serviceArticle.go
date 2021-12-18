package service

import (
	"github.com/luenci/go-gin-example/models"
	"github.com/luenci/go-gin-example/pkg/http/paginate"
	"github.com/luenci/go-gin-example/types/request"
)

type articleService struct {
}

func (a *articleService) ListArticleService(query *paginate.Query) (*models.ArticleList, error) {
}

func (a *articleService) GetArticleService(id uint) (*models.Article, error) {
}

func (a *articleService) DeleteArticleService(id uint) (*models.Article, error) {
}

func (a *articleService) CreateArticleService(request.CreateArticleRequest) (*models.Article, error) {
}

func (a *articleService) UpdateArticleService(request.UpdateArticleRequest) (*models.Article, error) {
}
