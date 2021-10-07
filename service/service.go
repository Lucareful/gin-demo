package service

import (
	"github.com/luenci/go-gin-example/types/request"
	"github.com/luenci/go-gin-example/types/response"
)

type Service struct {
	Tag TagService
}

type TagService interface {
	ListTagService(request.ListTagRequest) response.ListTagResponse
	CreateTagService(request.CreateTagRequest) response.CreateTagResponse
	UpdateTagService(request.UpdateTagRequest) response.UpdateTagResponse
	DeleteTagService(request.DeleteTagRequest) response.DeleteTagResponse
}

func NewService() *Service {
	// 检查struct是否实现了接口。原理:具体类型转换成接口类型，假如实现了接口类型，编译就可以通过，没有现实编译就会出错
	return &Service{Tag: (*Tagservice)(nil)}
}