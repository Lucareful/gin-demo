package render

import (
	"sync"

	"github.com/luenci/go-gin-example/pkg/e"

	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
)

type PaginateResult struct {
	Data *PaginateData `json:"data"` // 返回内容
	Msg  string        `json:"msg"`  // 请求结果
	Code int           `json:"code"` // 状态码
}

type PaginateData struct {
	Items    interface{}   `json:"items"`
	Paginate *PaginateInfo `json:"paginate"`
}

type PaginateInfo struct {
	Total    int64 `json:"total"`     // 数据总数
	Page     int   `json:"page"`      // 页数
	PageSize int   `json:"page_size"` // 每页数据量
}

func (p *PaginateResult) reset() {
	p.Data = nil
	p.Msg = ""
	p.Code = 0
}

var paginateResultPool = &sync.Pool{
	New: func() interface{} {
		return new(PaginateResult)
	},
}

func PaginateResponse(ctx *gin.Context, count int64, page, pageSize int, data ...interface{}) {
	result := paginateResultPool.Get().(*PaginateResult)
	defer func() {
		result.reset()
		paginateResultPool.Put(result)
	}()
	var paginateInfo PaginateInfo

	if pageSize == 0 {
		pageSize = int(count)
	}
	paginateInfo.Total = count
	paginateInfo.Page = page
	paginateInfo.PageSize = pageSize

	var err error
	if len(data) >= 1 {
		if _, ok := data[0].(error); ok {
			err = data[0].(error)
		}
	}
	var httpStatus int
	if err != nil {
		coder := errors.ParseCoder(err)
		result.Code = coder.Code()
		result.Msg = coder.String()
		httpStatus = coder.HTTPStatus()
	} else {
		result.Code = e.SUCCESS
		result.Msg = "Success"
		httpStatus = 200
	}

	var paginateData PaginateData
	paginateData.Paginate = &paginateInfo
	switch {
	case httpStatus >= 400 && httpStatus < 500:
		ctx.Abort()
		ctx.Set("warn", err)
		paginateData.Items = err.Error()
	case httpStatus >= 500:
		ctx.Abort()
		ctx.Set("error", err)
	default:
		if len(data) >= 1 {
			paginateData.Items = data[0]
		}
	}
	result.Data = &paginateData
	ctx.JSON(httpStatus, result)
}
