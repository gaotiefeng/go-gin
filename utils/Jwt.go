package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

const (
	JWT_TOKEN_TIME int = 3600
	JWT_SECRET  string = "JWT-ARY-STARK"
)
type JWTClaims struct { // token里面添加用户信息，验证token后可能会用到用户信息
	jwt.StandardClaims
	UserID      int      `json:"user_id"`
	Password    string   `json:"password"`
	Username    string   `json:"username"`
	FullName    string   `json:"full_name"`
	Permissions []string `json:"permissions"`
}

func GetToken(claims *JWTClaims)(string,error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return "",errors.New(JWT_SECRET)
	}
	return signedToken,nil
}

