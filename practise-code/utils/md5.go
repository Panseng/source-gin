package utils

import (
	"crypto/md5"
	"encoding/hex"
	"practise-code/global"
)

func MD5(pwd string) string{
	b := []byte(pwd)
	h := md5.New()
	h.Write(b)
	// 加盐
	h.Write([]byte(global.CONFIG.Salt))
	return hex.EncodeToString(h.Sum(nil))
}
