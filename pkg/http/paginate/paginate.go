package paginate

import (
	"strings"

	"github.com/luenci/go-gin-example/pkg/joint"

	"github.com/wxnacy/wgo/arrays"
	"gorm.io/gorm"
)

var IgnoreParams = []string{"search", "page", "order", "order_by", "page_size"}

const (
	DefaultPageSizeStr = "20"
	DefaultPageStr     = "1"
	DefaultPageSize    = 20
	DefaultPage        = 1
	DefaultSearch      = ""
	DefaultOrder       = "desc"
	DefaultOrderBy     = "id"
)

// Query 分页查询
type Query struct {
	Page     int                 `default:"1" minimum:"-1"`                                                   // 页码
	PageSize int                 `default:"20" minimum:"-1" maximum:"1000" json:"page_size" form:"page_size"` // 每页大小(最大1000条)
	Order    string              `default:"desc" enums:"desc,asc"`                                            // 排序参数(desc:倒序、asc：正序)
	OrderBy  string              `default:"id" json:"order_by" form:"order_by"`                               // 排序对象
	Search   string              // 模糊查询
	Params   map[string][]string `swaggerignore:"true"`
	IDList   []uint              `swaggerignore:"true"`
	AllData  bool                `json:"all_data" form:"all_data"`
}

func Paginate(q *Query, fuzzySearchFieldList, fieldList []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// 分页
		if !q.AllData {
			db = db.Offset((q.Page - 1) * q.PageSize).Limit(q.PageSize)
		}
		// 精确查询
		for k, v := range q.Params {
			if arrays.Contains(fieldList, k) != -1 {
				if len(v) == 1 {
					if v[0] != "" {
						db = db.Where(joint.ConcatStr(k, " = ?"), v[0])
					}
				} else {
					db = db.Where(joint.ConcatStr(k, " in (?)"), v)
				}
			}
		}

		// 模糊查询
		if q.Search != "" && len(fuzzySearchFieldList) > 0 {
			var searchField = make([]string, 0, len(fuzzySearchFieldList))
			for i := range fuzzySearchFieldList {
				searchField = append(searchField, joint.ConcatStr("IFNULL(`", strings.TrimSpace(fuzzySearchFieldList[i]), "`, '')"))
			}
			db = db.Where(joint.ConcatStr("concat(", strings.Join(searchField, ", "), ") like ?"), joint.ConcatStr("%", q.Search, "%"))
		}

		// 排序
		if arrays.Contains(fieldList, q.OrderBy) == -1 {
			q.OrderBy = DefaultOrderBy
		}

		if arrays.Contains([]string{"desc", "asc"}, q.Order) == -1 {
			q.Order = DefaultOrder
		}

		db = db.Order(joint.ConcatStr(q.OrderBy, " ", q.Order))

		// ID列表筛选
		if len(q.IDList) > 0 {
			db = db.Where("id in (?)", q.IDList)
		}

		return db
	}
}
