package v1

import (
	"github.com/luenci/go-gin-example/pkg/http/render"
	"github.com/marmotedu/errors"

	"github.com/gin-gonic/gin"
	"github.com/luenci/go-gin-example/pkg/e"
	http "github.com/luenci/go-gin-example/pkg/http/bind"
	"github.com/luenci/go-gin-example/types/request"
)

// ListTag godoc
// @Summary ListTag all tag - 返回所有 tag 数据.
// @Description ListTag all tag
// @Tags tag
// @Accept  json
// @Produce  json
// @Param tag body paginate.Query true "分页查询对象"
// @Success 200 {object} render.PaginateResult{data=[]models.Tag} "请求成功"
// @Failure 500 {object} render.HTTPError
// @Router /api/v1/tag [get]
func ListTag(c *gin.Context) {
	query := http.PaginateQuery(c)

	response, err := svc.Tag.ListTagService(query)

	if err != nil {
		render.PaginateResponse(c, 0, 0, 0, err)
		return
	}

	render.PaginateResponse(c, response.TotalCount, query.Page, query.PageSize, response.Item)

}

// CreateTag godoc
// @Summary create a tag - 创建一个文章标签
// @Description create an article tag
// @Tags tag
// @Accept  json
// @Produce  json
// @Param tag body request.CreateTagRequest true "Add article tag"
// @Success 200 {object} response.CreateTagResponse
// @Failure 500 {object} render.HTTPError
// @Router /api/v1/tag [post]
func CreateTag(c *gin.Context) {
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

// UpdateTag godoc
// @Summary UpdateTag a tag - 更新文章标签
// @Description UpdateTag an article tag
// @Tags tag
// @Accept  json
// @Produce  json
// @Param  tag body request.UpdateTagRequest true "UpdateTag article tag"
// @Success 200 {object} response.UpdateTagResponse
// @Failure 500 {object} render.HTTPError
// @Router /api/v1/tag [patch]
func UpdateTag(c *gin.Context) {
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

// DeleteTag godoc
// @Summary DeleteTag a tag - 删除文章标签
// @Description DeleteTag an article tag
// @Tags tag
// @Accept  json
// @Produce  json
// @Param  id path int true "Tag ID" Format(uint)
// @Success 200 {object} render.HTTPSuccess
// @Failure 500 {object} render.HTTPError
// @Router /api/v1/tag/{id} [delete]
func DeleteTag(c *gin.Context) {

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

// GetTag godoc
// @Summary GetTag a tag - 获取单个文章标签
// @Description GetTag an article tag (获取单个文章标签)
// @Tags tag
// @Accept  json
// @Produce  json
// @Param  id path int true "Tag ID" Format(uint)
// @Success 200 {object} response.GetTagResponse
// @Failure 500 {object} render.HTTPError
// @Router /api/v1/tag/{id} [get]
func GetTag(c *gin.Context) {

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
