package jwts

import (
	"gvd_server/global"
	"gvd_server/models"

	"github.com/dgrijalva/jwt-go/v4"
)

// 载荷
type JwtPayLoad struct {
	NickName string `json:"nickName"`
	RoleID   uint   `json:"roleID"`
	UserID   uint   `json:"userID"`
	UserName string `json:"userName"`

}

// 密钥
var Secret []byte

// 自定义声明，用于描述 Token 信息的部分，比如过期时间、签发者、接收者等标准声明
// 用于生成最终jwt token
type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims //标准声明
}

func (c CustomClaims) GetUser() (user *models.UserModel, err error) {
	user = new(models.UserModel)
	err = global.DB.Take(user, c.UserID).Error
	return
}
