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

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Study Swagger API"
	docs.SwaggerInfo.Description = "This is a sample server API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "127.0.0.1:8000"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiv1.GET("/tag", v1.ListTag)

		apiv1.GET("/tag/:id", v1.GetTag)
		//新建标签
		apiv1.POST("/tag", v1.CreateTag)
		//更新指定标签
		apiv1.PATCH("/tag", v1.UpdateTag)
		//删除指定标签
		apiv1.DELETE("/tag/:id", v1.DeleteTag)

		//获取用户列表
		apiv1.GET("/user", v1.ListUser)
		// 获取指定用户
		apiv1.GET("/user/:id", v1.GetUser)
		// 新建用户
		apiv1.POST("/user", v1.CreateUser)
		// 更新指定用户
		apiv1.PATCH("/user/:id", v1.UpdateUser)
		// 删除指定用户
		apiv1.DELETE("/user/:id", v1.DeleteUser)
	}

	return r
}
