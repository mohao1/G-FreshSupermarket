package Utile

import (
	"crypto/md5"
	"fmt"
)

// StrMD5ByStr String=>MD5-String
func StrMD5ByStr(str string) string {
	return ByteMD5Str([]byte(str))
}

// ByteMD5Str :Byte=>MD5-String
func ByteMD5Str(b []byte) string {
	m := md5.New()
	_, err := m.Write(b)
	if err != nil {
		panic(err)
	}
	sum := m.Sum([]byte(md5KEY))
	return fmt.Sprintf("%x", sum)
}
