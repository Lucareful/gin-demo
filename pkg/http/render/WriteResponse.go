package render

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
)

// Response 不分页的返回响应结构体
type Response struct {
	Code    int                    `json:"code,omitempty"`
	Message string                 `json:"message,omitempty"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

// WriteResponse 返回响应的数据.
func WriteResponse(c *gin.Context, data ...interface{}) {
	if err != nil {
		coder := errors.ParseCoder(err)
		c.JSON(coder.HTTPStatus(), Response{
			Code:    coder.Code(),
			Message: coder.String(),
			Data:    data,
		})
	}

	c.JSON(http.StatusOK, Response{Data: data})
}
