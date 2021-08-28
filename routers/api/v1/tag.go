package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	models "github.com/luenci/go-gin-example/models"
	"github.com/luenci/go-gin-example/pkg/e"
	"github.com/luenci/go-gin-example/service"
	"github.com/luenci/go-gin-example/types/request"
	"github.com/unknwon/com"
)

// List 获取多个文章标签
func List(c *gin.Context) {
	var r request.ListTagRequest
	var q service.Tagservice

	if err := c.ShouldBindQuery(&r); err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	response := q.ListTagService(r)

	c.JSON(http.StatusOK, response)
}

// Create 新增文章标签
func Create(c *gin.Context) {
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

// Update 修改文章标签
func Update(c *gin.Context) {
}

// Delete 删除文章标签
func Delete(c *gin.Context) {
}
