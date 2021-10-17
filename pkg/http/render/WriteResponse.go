package render

import (
	"sync"

	"github.com/luenci/go-gin-example/pkg/e"

	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
)

// Response 不分页的返回响应结构体
type Response struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (r *Response) reset() {
	r.Data = nil
	r.Message = ""
	r.Code = 0
}

// 请求结构体复用
var responsePool = &sync.Pool{
	New: func() interface{} {
		return new(Response)
	},
}

// WriteResponse 返回响应的数据.
func WriteResponse(c *gin.Context, err error, data interface{}) {

	response := responsePool.Get().(*Response)
	defer func() {
		response.reset()
		responsePool.Put(response)
	}()

	var httpStatus int

	if err != nil {
		coder := errors.ParseCoder(err)
		response.Code = coder.Code()
		response.Message = coder.String()
		httpStatus = coder.HTTPStatus()
	} else {
		response.Code = e.SUCCESS
		response.Message = "Success"
		httpStatus = 200
	}

	switch {
	case httpStatus >= 400 && httpStatus < 500:
		c.Abort()
		c.Set("warn", err)
		response.Data = err.Error()
	case httpStatus >= 500:
		c.Abort()
		c.Set("error", err)
	default:
		response.Data = data
	}
	c.JSON(httpStatus, response)
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"Service internal error"`
}

type HTTPSuccess struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"request ok 200"`
}
