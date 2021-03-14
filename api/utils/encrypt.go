package utils

import (
	"crypto/sha1"
	"fmt"
)

// Hashing "SHA-1"を使用して160ビットのハッシュ値を生成
func Hashing(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
