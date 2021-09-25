package service

import (
	"github.com/luenci/go-gin-example/models"
	"github.com/luenci/go-gin-example/pkg/e"
	"github.com/luenci/go-gin-example/pkg/setting"
	"github.com/luenci/go-gin-example/types/request"
	"github.com/luenci/go-gin-example/types/response"
)

type Tagservice struct{}

func (t *Tagservice) ListTagService(r request.ListTagRequest) response.ListTagResponse {

	var q response.ListTagResponse

	q.Data = make(map[string]interface{})

	q.Data["state"] = r.State

	q.Code = e.SUCCESS

	// todo: 分页
	q.Data["lists"] = models.GetTags(1, setting.PageSize, q.Data["state"])
	q.Data["total"] = models.GetTagTotal(q.Data["state"])

	q.Msg = e.GetMsg(q.Code)

	return q
}

func (t *Tagservice) CreateTagService(r request.CreateTagRequest) response.CreateTagResponse {
	return response.CreateTagResponse{}
}

func (t *Tagservice) UpdateTagService(r request.UpdateTagRequest) response.UpdateTagResponse {
	return response.UpdateTagResponse{}
}

func (t *Tagservice) DeleteTagService(r request.DeleteTagRequest) response.DeleteTagResponse {
	return response.DeleteTagResponse{}
}
