package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/luenci/go-gin-example/models"
)

func Login(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	// 测试
	if user.UserName != "luenci" && user.Password != "123456" {
		c.JSON(400, gin.H{
			"info": "用户未登陆",
		})
		return
	}

	tokenString, err := models.GenerateToken(user.UserName, user.Password)

	fmt.Println(tokenString, err)

	c.SetCookie("token", tokenString, -1, "/", "", false, true)
	c.Header("Access-Control-Allow-Origin", "*")

	c.JSON(200, "登陆成功")
}
