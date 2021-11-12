package service

import (
	"github.com/luenci/go-gin-example/models"
	"github.com/luenci/go-gin-example/pkg/http/paginate"
	"github.com/luenci/go-gin-example/types/request"
)

type tagService struct{}

func (t *tagService) ListTagService(query *paginate.Query) (*models.TagList, error) {
	listTag := models.NewListTag()

	err := listTag.ListTag(query)
	if err != nil {
		return nil, err
	}

	return listTag, nil
}

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
