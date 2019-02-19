package cryptoutils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// MD5 md5加密
func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// SHA256 SHA256加密
func SHA256(passwd string) string {
	h := sha256.New()
	h.Write([]byte(passwd))
	return fmt.Sprintf("%x", h.Sum(nil))
}
