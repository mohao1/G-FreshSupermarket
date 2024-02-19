package Utile

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func GetJWTToken(userId string) (string, error) {
	claims := make(jwt.MapClaims)
	//签发时间
	claims["iat"] = time.Now().Unix()
	//用户id
	claims["jwtUserId"] = userId
	fmt.Println(claims)
	//编码
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	stk, _ := token.SignedString([]byte(tokenKEY))
	fmt.Println(stk)
	return stk, nil
}
