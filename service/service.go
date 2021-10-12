package service

import (
	"github.com/luenci/go-gin-example/types/request"
	"github.com/luenci/go-gin-example/types/response"
)

type Service struct {
	Tag TagService
}

type TagService interface {
	//ListTagService(request.ListTagRequest) (*response.ListTagResponse, error)
	GetTagService(request.GetTagsRequest) (*response.GetTagResponse, error)
	CreateTagService(request.CreateTagRequest) (*response.CreateTagResponse, error)
	UpdateTagService(request.UpdateTagRequest) (*response.UpdateTagResponse, error)
	DeleteTagService(request.DeleteTagRequest) (*response.DeleteTagResponse, error)
}

func NewService() *Service {
	// 检查struct是否实现了接口。原理:具体类型转换成接口类型，假如实现了接口类型，编译就可以通过，没有现实编译就会出错
	return &Service{Tag: (*tagService)(nil)}
}
