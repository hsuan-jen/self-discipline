package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//字符串切片合并为单个字符串
func MergeStr(strs []interface{}) string {
	var build strings.Builder

	for _, str := range strs {
		build.WriteString(fmt.Sprintf("%v", str))
	}

	return build.String()
}

//随机字符串
func RandomStr(way uint8, l int) string {

	var bytes []byte
	switch way {
	case 1:
		bytes = []byte("0123456789")
	case 2:
		bytes = []byte("0123456789abcdefghijklmnopqrstuvwxyz")
	default:
		bytes = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	}

	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)

}
