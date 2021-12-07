package v1

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/luenci/go-gin-example/models"
)

func Login(c *gin.Context) {
	var user models.User
	var hmacSampleSecret []byte

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	if user.UserName != "luenci" && user.Password != "123456" {
		c.JSON(400, gin.H{
			"info": "用户未登陆",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		user.UserName: user.Password,
		"time":        time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)

	fmt.Println(tokenString, err)

	c.SetCookie("token", "", -1, "/", "", false, true)
	c.JSON(200, "登陆成功")
}
