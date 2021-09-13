package tests

import (
	"fmt"
	"testing"

	util "github.com/luenci/go-gin-example/pkg/util"
)

func TestConcatStr(t *testing.T) {
	if str := util.ConcatStr("123", "456", "789"); str != "123456789" {
		t.Errorf("export 123456789, but %s got", str)
	}
	fmt.Println("success")
}
