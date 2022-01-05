package middleware

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/casbin/casbin/v2"
)

// ValidAuth 权限校验
func ValidAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		e, err := casbin.NewEnforcer("../pkg/rbac/model.config", "../pkg/rbac/policy.csv")
		if err != nil {
		}
		roles, _ := e.GetRolesForUser("alice")
		fmt.Println(roles)

		ok, _ := e.Enforce("alice", "data1", "read")
		if !ok {
			log.Fatal("权限校验失败")
		}

		c.Next()
	}

}
