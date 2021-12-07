package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/luenci/go-gin-example/pkg/e"
	http "github.com/luenci/go-gin-example/pkg/http/bind"
	"github.com/luenci/go-gin-example/pkg/http/render"
	"github.com/luenci/go-gin-example/types/request"
	"github.com/marmotedu/errors"
)

// ListUser godoc
// @Summary ListUser all user - 返回所有 user 数据.
// @Description ListUser all user
// @Tags tag
// @Accept  json
// @Produce  json
// @Param tag body paginate.Query true "分页查询对象"
// @Success 200 {object} render.PaginateResult{data=[]models.User} "请求成功"
// @Failure 500 {object} render.HTTPError
// @Router /api/v1/user [get]
func ListUser(c *gin.Context) {
	query := http.PaginateQuery(c)

	response, err := svc.User.ListUserService(query)

	if err != nil {
		render.PaginateResponse(c, 0, 0, 0, err)
		return
	}

	render.PaginateResponse(c, response.TotalCount, query.Page, query.PageSize, response.Item)

}

// CreateUser godoc
// @Summary create a user - 创建一个用户标签
// @Description create a user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body request.CreateUserRequest true "Add article user"
// @Success 200 {object} models.User "请求成功"
// @Failure 500 {object} render.HTTPError
// @Router /api/v1/user [post]
func CreateUser(c *gin.Context) {
	var r request.CreateUserRequest

	if err := http.Body(c, &r); err != nil {
		return
	}

	response, err := svc.User.CreateUserService(r)
	if err != nil {
		render.WriteResponse(c, errors.WithCode(e.ERROR_EXIST_TAG, e.GetMsg(e.ERROR_EXIST_TAG)), err)
		return
	}
	render.WriteResponse(c, errors.WithCode(e.SUCCESS, e.GetMsg(e.SUCCESS)), response)

}

// UpdateUser godoc
// @Summary UpdateUser a user - 更新用户信息
// @Description UpdateUser a user info
// @Tags user
// @Accept  json
// @Produce  json
// @Param  user body request.UpdateUserRequest true "UpdateTag article user"
// @Success 200 {object} models.User
// @Failure 500 {object} render.HTTPError
// @Router /api/v1/user [patch]
func UpdateUser(c *gin.Context) {
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

// DeleteUser godoc
// @Summary DeleteUser a user - 删除文章标签
// @Description DeleteUser an article user
// @Tags user
// @Accept  json
// @Produce  json
// @Param  id path int true "Tag ID" Format(uint)
// @Success 200 {object} render.HTTPSuccess
// @Failure 500 {object} render.HTTPError
// @Router /api/v1/user/{id} [delete]
func DeleteUser(c *gin.Context) {

	id, err := http.ParamsID(c, "id")
	if err != nil {
		return
	}

	response, err := svc.User.DeleteUserService(id)
	if err != nil {
		render.WriteResponse(c, errors.WithCode(e.ERROR_NOT_EXIST_TAG, e.GetMsg(e.ERROR_NOT_EXIST_TAG)), nil)
		return
	}

	render.WriteResponse(c, errors.WithCode(e.SUCCESS, e.GetMsg(e.SUCCESS)), response)
}

// GetUser godoc
// @Summary GetUser a user - 获取用户信息.
// @Description GetUser a user (获取单个用户)
// @Tags user
// @Accept  json
// @Produce  json
// @Param  id path int true "Tag ID" Format(uint)
// @Success 200 {object} models.User
// @Failure 500 {object} render.HTTPError
// @Router /api/v1/user/{id} [get]
func GetUser(c *gin.Context) {

	id, err := http.ParamsID(c, "id")
	if err != nil {
		return
	}

	response, err := svc.User.GetUserService(id)
	if err != nil {
		render.WriteResponse(c, errors.WithCode(e.ERROR_NOT_EXIST_TAG, e.GetMsg(e.ERROR_NOT_EXIST_TAG)), nil)
		return
	}

	render.WriteResponse(c, errors.WithCode(e.SUCCESS, e.GetMsg(e.SUCCESS)), response)
}
