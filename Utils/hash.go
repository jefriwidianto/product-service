package Utils

import (
	"crypto/md5"
	"fmt"
)

func GenerateHashPass(secret string) string {
	data := []byte(secret)
	hash := fmt.Sprintf("%x", md5.Sum(data))
	return hash
}
