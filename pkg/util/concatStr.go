package util

import "strings"

// ConcatStr 字符串拼接
func ConcatStr(str ...string) string {
	var builder strings.Builder
	builder.Grow(len(str))
	for _, v := range str {
		builder.WriteString(v)
	}

	return builder.String()
}
