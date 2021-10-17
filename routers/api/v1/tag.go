package v1

import (
	"github.com/luenci/go-gin-example/pkg/http/render"

	"github.com/marmotedu/errors"

	"github.com/gin-gonic/gin"
	"github.com/luenci/go-gin-example/pkg/e"
	"github.com/luenci/go-gin-example/types/request"
)

// List 获取多个文章标签
//func List(c *gin.Context) {
//	var r request.ListTagRequest
//
//	if err := c.ShouldBindQuery(&r); err != nil {
//		erros := make(map[string]interface{})
//		erros["code"] = e.VALIDARION_ERRORS
//		erros["message"] = e.GetMsg(e.VALIDARION_ERRORS)
//
//		c.AbortWithStatusJSON(
//			http.StatusInternalServerError,
//			gin.H{"data": erros})
//		return
//	}
//
//	//response, err := svc.Tag.ListTagsService(r)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, errors.WithCode(e.ERROR_NOT_EXIST_TAG, e.GetMsg(e.ERROR_NOT_EXIST_TAG)))
//		return
//	}
//
//	c.JSON(http.StatusOK, response)
//}

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

	if err := c.ShouldBindJSON(&r); err != nil {
		render.WriteResponse(c, errors.WithCode(e.VALIDARION_ERRORS, e.GetMsg(e.VALIDARION_ERRORS)), nil)
		return
	}

	response, err := svc.Tag.CreateTagService(r)
	if err != nil {
		render.WriteResponse(c, errors.WithCode(e.ERROR_EXIST_TAG, e.GetMsg(e.ERROR_EXIST_TAG)), nil)
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

	if err := c.ShouldBindJSON(&r); err != nil {
		render.WriteResponse(c, errors.WithCode(e.VALIDARION_ERRORS, e.GetMsg(e.VALIDARION_ERRORS)), nil)
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
	var id uint

	if err := c.ShouldBindQuery(&id); err != nil {
		render.WriteResponse(c, errors.WithCode(e.ERROR_EXIST_TAG, e.GetMsg(e.ERROR_EXIST_TAG)), nil)
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
	var id uint

	if err := c.ShouldBindQuery(&id); err != nil {
		render.WriteResponse(c, errors.WithCode(e.ERROR_EXIST_TAG, e.GetMsg(e.ERROR_EXIST_TAG)), nil)
		return
	}

	response, err := svc.Tag.GetTagService(id)
	if err != nil {
		render.WriteResponse(c, errors.WithCode(e.ERROR_NOT_EXIST_TAG, e.GetMsg(e.ERROR_NOT_EXIST_TAG)), nil)
		return
	}

	render.WriteResponse(c, errors.WithCode(e.SUCCESS, e.GetMsg(e.SUCCESS)), response)
}
