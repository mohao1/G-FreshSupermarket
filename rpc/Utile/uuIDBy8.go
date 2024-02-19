package Utile

import "github.com/google/uuid"

func GetUUIDBy8() string {
	uuidWithHyphen := uuid.New().String()
	uuidWithoutHyphenStr := uuidWithoutHyphen(uuidWithHyphen)
	return uuidWithoutHyphenStr[:8]
}

func uuidWithoutHyphen(uuidWithHyphen string) string {
	return uuidWithHyphen[0:8] +
		uuidWithHyphen[9:13] +
		uuidWithHyphen[14:18] +
		uuidWithHyphen[19:23] +
		uuidWithHyphen[24:]
}
