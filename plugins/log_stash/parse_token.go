package log_stash

import (

	"github.com/dgrijalva/jwt-go/v4"
)

type JwtPayLoad struct {
	NickName string `json:"nickName"`
	RoleID   uint   `json:"roleID"`
	UserID   uint   `json:"userID"`
	UserName string `json:"userName"`
}

// 自定义声明，用于描述 Token 信息的部分，比如过期时间、签发者、接收者等标准声明
// 用于生成最终jwt token
type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims //标准声明
}

func parseToken(token string) (jwtPayload *JwtPayLoad) {
	Token, _ := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})

	//Token == nil 检查传入的token是否为空
	//Token.Claims == nil 检查传入的token是否完整，解析出的claims是否有效
	if Token == nil || Token.Claims == nil {
		return nil
	}
	claims, ok := Token.Claims.(*CustomClaims)
	if !ok {
		return nil
	}
	return &claims.JwtPayLoad
}