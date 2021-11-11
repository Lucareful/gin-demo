package service

import (
	"github.com/luenci/go-gin-example/models"
	"github.com/luenci/go-gin-example/types/request"
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

func (t *tagService) GetTagService(id uint) (*models.Tag, error) {
	tag := models.NewTag()

	if err := tag.GetTag(id); err != nil {
		return nil, err
	}

	return tag, nil
}

func (t *tagService) CreateTagService(r request.CreateTagRequest) (*models.Tag, error) {

	tag := models.NewTag()

	if err := tag.CreateTag(r); err != nil {
		return nil, err
	}
	tag.Name = r.Name
	tag.CreatedBy = r.CreatedBy
	tag.State = r.State

	return tag, nil
}

func (t *tagService) UpdateTagService(r request.UpdateTagRequest) (*models.Tag, error) {
	tag := models.NewTag()

	if err := tag.UpdateTag(r); err != nil {
		return nil, err
	}

	return tag, nil
}

func (t *tagService) DeleteTagService(id uint) (*models.Tag, error) {
	tag := models.NewTag()

	if err := tag.DeleteTag(id); err != nil {
		return nil, err
	}

	return tag, nil

}
