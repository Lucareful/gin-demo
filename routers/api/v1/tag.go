package v1

import (
	"github.com/luenci/go-gin-example/pkg/http/render"
	"github.com/marmotedu/errors"

	"github.com/gin-gonic/gin"
	"github.com/luenci/go-gin-example/pkg/e"
	http "github.com/luenci/go-gin-example/pkg/http/bind"
	"github.com/luenci/go-gin-example/types/request"
)

// List godoc
// @Summary List all tag - 返回所有 tag 数据.
// @Description List all tag
// @Tags tag
// @Accept  json
// @Produce  json
// @Param tag body paginate.Query true "List all tag"
// @Success 200 {object} render.PaginateResponse
// @Failure 500 {object} render.HTTPError
// @Router /api/v1/tag [list]
func List(c *gin.Context) {
	query := http.PaginateQuery(c)

	response, err := svc.Tag.ListTagService(query)

	if err != nil {
		render.PaginateResponse(c, 0, 0, 0, err)
		return
	}

	render.PaginateResponse(c, response.TotalCount, query.Page, query.PageSize, response.Item)

}

// Create godoc
// @Summary create a tag - 创建一个文章标签
// @Description create an article tag
// @Tags tag
// @Accept  json
// @Produce  json
// @Param tag body request.CreateTagRequest true "Add article tag"
// @Success 200 {object} response.CreateTagResponse
// @Failure 500 {object} render.HTTPError
// @Router /api/v1/tag [post]
func Create(c *gin.Context) {
	var r request.CreateTagRequest

	if err := http.Body(c, &r); err != nil {
		return
	}

	response, err := svc.Tag.CreateTagService(r)
	if err != nil {
		render.WriteResponse(c, errors.WithCode(e.ERROR_EXIST_TAG, e.GetMsg(e.ERROR_EXIST_TAG)), err)
		return
	}
	render.WriteResponse(c, errors.WithCode(e.SUCCESS, e.GetMsg(e.SUCCESS)), response)

}

// Update godoc
// @Summary Update a tag - 更新文章标签
// @Description Update an article tag
// @Tags tag
// @Accept  json
// @Produce  json
// @Param  tag body request.UpdateTagRequest true "Update article tag"
// @Success 200 {object} response.UpdateTagResponse
// @Failure 500 {object} render.HTTPError
// @Router /api/v1/tag [patch]
func Update(c *gin.Context) {
	var r request.UpdateTagRequest

	if err := http.Body(c, &r); err != nil {
		return
	}

	response, err := svc.Tag.UpdateTagService(r)
	if err != nil {
		render.WriteResponse(c, errors.WithCode(e.ERROR_EXIST_TAG, e.GetMsg(e.ERROR_EXIST_TAG)), nil)
		return
	}

	render.WriteResponse(c, errors.WithCode(e.SUCCESS, e.GetMsg(e.SUCCESS)), response)
}

// Delete godoc
// @Summary Delete a tag - 删除文章标签
// @Description Delete an article tag
// @Tags tag
// @Accept  json
// @Produce  json
// @Param  id path int true "Tag ID" Format(uint)
// @Success 200 {object} render.HTTPSuccess
// @Failure 500 {object} render.HTTPError
// @Router /api/v1/tag/{id} [delete]
func Delete(c *gin.Context) {

	id, err := http.ParamsID(c, "id")
	if err != nil {
		return
	}

	response, err := svc.Tag.DeleteTagService(id)
	if err != nil {
		render.WriteResponse(c, errors.WithCode(e.ERROR_NOT_EXIST_TAG, e.GetMsg(e.ERROR_NOT_EXIST_TAG)), nil)
		return
	}

	render.WriteResponse(c, errors.WithCode(e.SUCCESS, e.GetMsg(e.SUCCESS)), response)
}

// Get godoc
// @Summary Get a tag - 获取单个文章标签
// @Description Get an article tag (获取单个文章标签)
// @Tags tag
// @Accept  json
// @Produce  json
// @Param  id path int true "Tag ID" Format(uint)
// @Success 200 {object} response.GetTagResponse
// @Failure 500 {object} render.HTTPError
// @Router /api/v1/tag/{id} [get]
func Get(c *gin.Context) {

	id, err := http.ParamsID(c, "id")
	if err != nil {
		return
	}

	response, err := svc.Tag.GetTagService(id)
	if err != nil {
		render.WriteResponse(c, errors.WithCode(e.ERROR_NOT_EXIST_TAG, e.GetMsg(e.ERROR_NOT_EXIST_TAG)), nil)
		return
	}

	render.WriteResponse(c, errors.WithCode(e.SUCCESS, e.GetMsg(e.SUCCESS)), response)
}
