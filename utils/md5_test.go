package utils

import "testing"

func TestMD5V(t *testing.T) {
	str := MD5V("123456")
	t.Log(str)
}
