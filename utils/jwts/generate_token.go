package jwts

import (
	"gvd_server/global"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

func GenToken(jwtpayload JwtPayLoad) (string, error) {
	claims := CustomClaims{
		JwtPayLoad: jwtpayload,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Duration(global.Config.Jwt.Expires) * time.Hour)),
			Issuer:    global.Config.Jwt.Issuer,
		},
	}
	//编码并加密,创建了一个包含自定义声明claims的Token对象
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//将令牌对象使用预先设置的jwtSecret密钥进行签名并生成字符串形式的JWT令牌
	token, err := tokenClaims.SignedString([]byte(global.Config.Jwt.Secret))

	return token, err
}
