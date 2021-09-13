package setting

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// GetPage 分页器
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * PageSize
	}

	return result
}
