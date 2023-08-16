package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Hash(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
