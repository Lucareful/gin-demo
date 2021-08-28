package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	models "github.com/luenci/go-gin-example/models"
	"github.com/luenci/go-gin-example/pkg/e"
	"github.com/luenci/go-gin-example/pkg/setting"
	"github.com/luenci/go-gin-example/pkg/util"
	"github.com/luenci/go-gin-example/types/request"
	"github.com/luenci/go-gin-example/types/response"
	"github.com/unknwon/com"
)

// GetTags 获取多个文章标签
func GetTags(c *gin.Context) {
	var r request.ListTagRequest
	q := &response.ListTagResponse{}

	if err := c.ShouldBindQuery(&r); err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	q.Data = make(map[string]interface{})

	q.Data["state"] = r.State

	q.Code = e.SUCCESS

	q.Data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, q.Data["state"])
	q.Data["total"] = models.GetTagTotal(q.Data["state"])

	q.Msg = e.GetMsg(q.Code)

	c.JSON(http.StatusOK, q)
}

// AddTag 新增文章标签
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")

	valid := validator.New()
	newTag := &models.Tag{Name: name, State: state, CreatedBy: createdBy}
	err := valid.Struct(newTag)

	var code int
	if err != nil {
		code = e.ERROR_EXIST_TAG
	} else {
		code = e.SUCCESS
		models.AddTag(name, state, createdBy)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}

// EditTag 修改文章标签
func EditTag(c *gin.Context) {
}

// DeleteTag 删除文章标签
func DeleteTag(c *gin.Context) {
}
