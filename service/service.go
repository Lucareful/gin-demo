package service

import (
	"github.com/luenci/go-gin-example/models"
	"github.com/luenci/go-gin-example/pkg/http/paginate"
	"github.com/luenci/go-gin-example/types/request"
	"github.com/luenci/go-gin-example/types/response"
)

// Service 所有接口服务的注册.
type Service struct {
	Tag  TagService
	User UserService
}

// TagService tag 服务接口.
type TagService interface {
	ListTagService(query *paginate.Query) (*models.TagList, error)
	GetTagService(id uint) (*models.Tag, error)
	DeleteTagService(id uint) (*models.Tag, error)
	CreateTagService(request.CreateTagRequest) (*models.Tag, error)
	UpdateTagService(request.UpdateTagRequest) (*models.Tag, error)
}

type UserService interface {
	ListUserService(query *paginate.Query) (*models.UserList, error)
	GetUserService(id uint) (*models.User, error)
	DeleteUserService(id uint) (*models.User, error)
	CreateUserService(request.CreateUserRequest) (*models.User, error)
	UpdateUserService(request.UpdateUserRequest) (*models.User, error)
}

// ArticleService 文件服务接口.
type ArticleService interface {
	GetArticleService(id uint) (*response.GetTagResponse, error)
	DeleteArticleService(id uint) (*response.DeleteTagResponse, error)
	CreateArticleService(request.CreateTagRequest) (*response.CreateTagResponse, error)
	UpdateArticleService(request.UpdateTagRequest) (*response.UpdateTagResponse, error)
}

func NewService() *Service {
	// 检查struct是否实现了接口。原理:具体类型转换成接口类型，假如实现了接口类型，编译就可以通过，没有现实编译就会出错
	return &Service{
		Tag:  (*tagService)(nil),
		User: (*userService)(nil),
	}
}
