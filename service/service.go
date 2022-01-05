package service

import (
	"github.com/luenci/go-gin-example/models"
	"github.com/luenci/go-gin-example/pkg/http/paginate"
	"github.com/luenci/go-gin-example/types/request"
)

// Service 所有接口服务的注册.
type Service struct {
	Tag     TagService
	User    UserService
	Article ArticleService
}

// TagService tag 服务接口.
type TagService interface {
	ListTagService(query *paginate.Query) (*models.TagList, error)
	GetTagService(id uint) (*models.Tag, error)
	DeleteTagService(id uint) (*models.Tag, error)
	CreateTagService(request.CreateTagRequest) (*models.Tag, error)
	UpdateTagService(request.UpdateTagRequest) (*models.Tag, error)
}

// UserService user 服务接口.
type UserService interface {
	ListUserService(query *paginate.Query) (*models.UserList, error)
	GetUserService(id uint) (*models.User, error)
	DeleteUserService(id uint) (*models.User, error)
	CreateUserService(request.CreateUserRequest) (*models.User, error)
	UpdateUserService(request.UpdateUserRequest) (*models.User, error)
}

// ArticleService 文章服务接口.
type ArticleService interface {
	ListArticleService(query *paginate.Query) (*models.ArticleList, error)
	GetArticleService(id uint) (*models.Article, error)
	DeleteArticleService(id uint) (*models.Article, error)
	CreateArticleService(request.CreateArticleRequest) (*models.Article, error)
	UpdateArticleService(request.UpdateArticleRequest) (*models.Article, error)
}

func NewService() *Service {
	// 检查struct是否实现了接口。原理:具体类型转换成接口类型，假如实现了接口类型，编译就可以通过，没有现实编译就会出错
	return &Service{
		Tag:     (*tagService)(nil),
		User:    (*userService)(nil),
		Article: (*articleService)(nil),
	}
}
