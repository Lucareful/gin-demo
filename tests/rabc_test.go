package tests

import (
	"testing"

	"github.com/luenci/go-gin-example/middleware"
)

func TestValidAuth(t *testing.T) {
	if ok := middleware.ValidAuth(); ok != true {
		t.Errorf("权限校验失败")
	}
}
