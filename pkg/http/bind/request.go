package bind

import (
	"bytes"
	"io"
	"strconv"

	ecode "github.com/luenci/go-gin-example/pkg/e"

	"github.com/luenci/go-gin-example/pkg/http/render"

	"github.com/marmotedu/errors"

	"github.com/gin-gonic/gin"
)

func Body(ctx *gin.Context, ptr interface{}) error {
	var err error
	if err = ctx.ShouldBind(ptr); err != nil {
		render.WriteResponse(ctx, errors.WithCode(ecode.BIND_PARAMS_FAIL, ecode.GetMsg(ecode.BIND_PARAMS_FAIL)), err)
	}
	return err
}

func Query(ctx *gin.Context, ptr interface{}) error {
	var err error
	if err = ctx.ShouldBindQuery(ptr); err != nil {
		render.WriteResponse(ctx, errors.WithCode(ecode.BIND_PARAMS_FAIL, ecode.GetMsg(ecode.BIND_PARAMS_FAIL)), err)
	}
	return err
}

func File(ctx *gin.Context, key string) (string, []byte, error) {
	f, fh, err := ctx.Request.FormFile(key)
	if err != nil {
		render.WriteResponse(ctx, errors.WithCode(ecode.BIND_PARAMS_FAIL, ecode.GetMsg(ecode.BIND_PARAMS_FAIL)), err)
		return "", nil, err
	}
	var content bytes.Buffer
	buf := make([]byte, 4096)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err != io.EOF {
				render.WriteResponse(ctx, errors.WithCode(ecode.BIND_PARAMS_FAIL, ecode.GetMsg(ecode.BIND_PARAMS_FAIL)), err)
				return "", nil, err
			}
			break
		}
		content.Grow(n)
		content.Write(buf[:n])
	}
	return fh.Filename, content.Bytes(), nil
}

func ParamsID(ctx *gin.Context, key string) (uint, error) {
	id, err := strconv.Atoi(ctx.Param(key))
	if id < 0 || err != nil {
		render.WriteResponse(ctx, errors.WithCode(ecode.BIND_PARAMS_FAIL, key+" is invalid"), err)
	}
	return uint(id), err
}

//func PaginateQuery(ctx *gin.Context) *paginate.Query {
//	search := ctx.DefaultQuery("search", paginate.DefaultSearch)
//	pageString := ctx.DefaultQuery("page", paginate.DefaultPageStr)
//	pageSizeString := ctx.DefaultQuery("page_size", paginate.DefaultPageSizeStr)
//	order := ctx.DefaultQuery("order", paginate.DefaultOrder)
//	orderBy := ctx.DefaultQuery("order_by", paginate.DefaultOrderBy)
//	allDataString := ctx.DefaultQuery("all_data", "")
//	var idList []uint
//	ids, isExist := ctx.Get("id_list")
//	if isExist {
//		if _, ok := ids.([]uint); ok {
//			idList = ids.([]uint)
//		}
//	}
//	q := ctx.Request.URL.Query()
//	params := make(map[string][]string, len(q))
//	if q != nil {
//		for k := range q {
//			if arrays.Contains(paginate.IgnoreParams, k) == -1 {
//				params[k] = q[k]
//			}
//		}
//	}
//
//	page, err := strconv.Atoi(pageString)
//	if err != nil {
//		page = paginate.DefaultPage
//	}
//
//	if page < 1 {
//		page = paginate.DefaultPage
//	}
//
//	pageSize, err := strconv.Atoi(pageSizeString)
//	if err != nil {
//		pageSize = paginate.DefaultPageSize
//	}
//
//	if pageSize < 1 {
//		pageSize = paginate.DefaultPageSize
//	}
//
//	var allData bool
//	if arrays.ContainsString([]string{"", "false", "False", "FALSE", "0"}, allDataString) == -1 {
//		allData = true
//		page = 1
//		pageSize = 0
//	}
//
//	return &paginate.Query{
//		Page:     page,
//		PageSize: pageSize,
//		Order:    order,
//		OrderBy:  orderBy,
//		Search:   search,
//		Params:   params,
//		IDList:   idList,
//		AllData:  allData,
//	}
//}
