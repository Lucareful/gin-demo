package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Timter 计算请求耗时 middleware
func Timter() gin.HandlerFunc {
	return func(c *gin.Context) {

		nowTime := time.Now()

		//请求处理
		c.Next()

		costTime := time.Since(nowTime)
		url := c.Request.URL.String()
		fmt.Printf("the request URL %s cost %v\n", url, costTime)
	}
}

// 1.中间件代码最后即使没有调用Next()方法，后续中间件及handlers也会执行；
// 2.如果在中间件函数的非结尾调用Next()方法当前中间件剩余代码会被暂停执行，会先去执行后续中间件及handlers，等这些handlers全部执行完以后程序控制权会回到当前中间件继续执行剩余代码；
// 3.如果想提前中止当前中间件的执行应该使用return退出而不是Next()方法；
// 4.如果想中断剩余中间件及handlers应该使用Abort方法，但需要注意当前中间件的剩余代码会继续执行。。
