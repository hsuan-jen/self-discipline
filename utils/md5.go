package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5V(encryptStr string) string {
	h := md5.New()
	h.Write([]byte(encryptStr))
	return hex.EncodeToString(h.Sum(nil))
}
