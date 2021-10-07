package v1

import (
	"net/http"

	"github.com/marmotedu/errors"

	"github.com/gin-gonic/gin"
	models "github.com/luenci/go-gin-example/models"
	"github.com/luenci/go-gin-example/pkg/e"
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

	if err := c.ShouldBindQuery(&r); err != nil {
		c.JSON(http.StatusBadRequest, errors.WithCode(e.VALIDARION_ERRORS, e.GetMsg(e.VALIDARION_ERRORS)))
		return
	}

	if err := models.AddTag(r); err != nil {
		c.JSON(http.StatusInternalServerError, errors.WithCode(e.ERROR_EXIST_TAG, e.GetMsg(e.ERROR_EXIST_TAG)))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
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
