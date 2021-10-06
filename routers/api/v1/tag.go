package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/luenci/go-gin-example/models"
	"github.com/luenci/go-gin-example/pkg/e"
	errors "github.com/luenci/go-gin-example/pkg/errors"
	"github.com/luenci/go-gin-example/types/request"
)

// List 获取多个文章标签
func List(c *gin.Context) {
	var r request.ListTagRequest

	if err := c.ShouldBindQuery(&r); err != nil {
		erros := make(map[string]interface{})
		erros["code"] = e.VALIDARION_ERRORS
		erros["message"] = e.GetMsg(e.VALIDARION_ERRORS)

		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"data": erros})
		return
	}

	response := svc.Tag.ListTagService(r)

	c.JSON(http.StatusOK, response)
}

// Create 新增文章标签
func Create(c *gin.Context) {
	var r request.CreateTagRequest
	var code int

	if err := c.ShouldBindQuery(&r); err != nil {
		errors.WithCode(e.ERROR_EXIST_TAG)
		return
	}

	code = e.SUCCESS
	if err := models.AddTag(r); err != nil {
		return
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

// Get 获取单个文章标签
func Get(c *gin.Context) {

}
