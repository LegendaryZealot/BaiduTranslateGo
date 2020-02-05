package md5

import (
	"crypto/md5"
	"encoding/hex"
)

// Encryption : encryption .
func Encryption(str string) (result string) {
	h := md5.New()
	h.Write([]byte(str))
	result = hex.EncodeToString(h.Sum(nil))
	return
}
