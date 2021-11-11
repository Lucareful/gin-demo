package joint

import "strings"

// ConcatStr 字符串拼接
func ConcatStr(str ...string) string {
	var (
		builder strings.Builder
		n       int
	)
	for i := 0; i < len(str); i++ {
		n += len(str[i])
	}
	builder.Grow(n)
	// 直接写值的时候，str发生复制
	// 索引取值可减少内存复制
	for i := range str {
		builder.WriteString(str[i])
	}
	return builder.String()
}
