package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(data string) string {
	rawByte := []byte(data)
	hashedByte := md5.Sum(rawByte)
	return hex.EncodeToString(hashedByte[:])
}
