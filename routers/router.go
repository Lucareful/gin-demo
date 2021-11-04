package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/luenci/go-gin-example/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	middleware "github.com/luenci/go-gin-example/middleware"
	"github.com/luenci/go-gin-example/pkg/setting"
	v1 "github.com/luenci/go-gin-example/routers/api/v1"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(middleware.Timter())

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	// programatically set swagger info
	docs.SwaggerInfo.Title = "Study Swagger API"
	docs.SwaggerInfo.Description = "This is a sample server API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "127.0.1.1"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	{
		//获取标签列表
		//apiv1.GET("/tags", v1.List)

		apiv1.GET("/tags/:id", v1.Get)
		//新建标签
		apiv1.POST("/tags", v1.Create)
		//更新指定标签
		apiv1.PUT("/tags", v1.Update)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.Delete)
	}

	return r
}
