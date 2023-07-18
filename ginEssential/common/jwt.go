package common

import (
	"ginEssential/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// 根据用户信息生成token
func ReleaseToken(user model.User) (string, error) {
	// 设置token有效时间,7天
	exirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exirationTime.Unix(), // 过期时间
			IssuedAt:  time.Now().Unix(),    // 发放时间
			Issuer:    "ginEssential",       // 发放人
			Subject:   "user token",         // 主题
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey) // 生成签名字符串

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	// 解析token,第三个参数是一个回调函数，用于在token过期后执行续签操作
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
