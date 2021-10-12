package service

import (
	"github.com/luenci/go-gin-example/models"
	"github.com/luenci/go-gin-example/types/request"
	"github.com/luenci/go-gin-example/types/response"
)

type tagService struct{}

//func (t *tagService) ListTagService(r request.ListTagRequest) (*response.ListTagResponse, error) {
//
//	var q response.ListTagResponse
//
//	q.Data = make(map[string]interface{})
//
//	q.Data["state"] = r.State
//
//	q.Code = e.SUCCESS
//
//	// todo: 分页
//	q.Data["lists"] = models.GetTags(1, setting.PageSize, q.Data["state"])
//	q.Data["total"] = models.GetTagTotal(q.Data["state"])
//
//	q.Msg = e.GetMsg(q.Code)
//
//	return &q, nil
//}

func (t *tagService) GetTagService(request.GetTagRequest) (*response.GetTagResponse, error) {
	return nil, nil
}

func (t *tagService) CreateTagService(r request.CreateTagRequest) (*response.CreateTagResponse, error) {

	if err := models.AddTag(r); err != nil {
		return nil, nil
	}

	q := &response.CreateTagResponse{
		Name:      r.Name,
		CreatedBy: r.CreatedBy,
		State:     r.State,
	}

	return q, nil
}

func (t *tagService) UpdateTagService(r request.UpdateTagRequest) (*response.UpdateTagResponse, error) {
	return nil, nil
}

func (t *tagService) DeleteTagService(r request.DeleteTagRequest) (*response.DeleteTagResponse, error) {
	return nil, nil
}
