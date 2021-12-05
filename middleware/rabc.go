package middleware

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
)

// ValidAuth 权限校验
func ValidAuth() bool {
	e, err := casbin.NewEnforcer("../pkg/rbac/model.config", "../pkg/rbac/policy.csv")
	if err != nil {
		return false
	}
	roles, _ := e.GetRolesForUser("alice")
	fmt.Println(roles)

	ok, _ := e.Enforce("alice", "data1", "read")
	if !ok {
		log.Fatal("权限校验失败")
		return false
	}
	return true
}
