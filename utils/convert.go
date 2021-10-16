package utils

import (
	"strings"
)

//字符串切片合并为单个字符串
func MergeStr(strs []string) string {
	var build strings.Builder

	for _, str := range strs {
		build.WriteString(str)
	}

	return build.String()
}
