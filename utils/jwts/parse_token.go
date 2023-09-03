package jwts

import (
	"gvd_server/global"

	"github.com/dgrijalva/jwt-go/v4"
)

func ParseToken(token string) (*CustomClaims, error) {
	//在 jwt.ParseWithClaims 函数内部，会根据提供的密钥 jwtSecret 对令牌进行验证，并在验证通过时将令牌的内容解码到 Claims 结构体指针中
	tokenClaims, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Config.Jwt.Secret), nil
	})

	//检查tokenClaims是否成功被解析
	//如果token不为空，那么可以解析出结果
	if tokenClaims != nil {
		//检查tokenClaims是否包含有效的声明，并将声明强制转换为*Claims类型的结构体指针claims
		if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	//如果token为空，直接返回nil
	return nil, err
}

/*
服务器解析 JWT Token 的原理主要涉及三个步骤：验证、解码和验证声明。

验证签名：
服务器首先需要验证 JWT Token 的签名，确保令牌的完整性和真实性。JWT Token 包含头部（Header）、载荷（Payload）和签名（Signature）三部分。签名是由头部和载荷部分使用特定的算法和密钥生成的。服务器使用相同的密钥和算法对头部和载荷部分进行签名，然后将生成的签名与 JWT Token 中的签名进行比对。如果两者匹配，说明令牌未被篡改，并且可以被信任。

解码令牌：
如果签名验证通过，服务器接下来需要对令牌进行解码。JWT Token 的头部和载荷部分都是经过 Base64 编码的 JSON 字符串，因此服务器需要先对令牌进行 Base64 解码，得到头部和载荷的原始 JSON 数据。

验证声明：
解码后的载荷部分包含了 JWT Token 中的声明（Claims），其中包含了令牌所携带的信息，如用户名、角色、过期时间等。服务器需要验证这些声明是否有效，比如验证令牌是否已过期、是否来自受信任的发行者等。验证声明时，服务器可能需要检查数据库或其他存储来获取与用户相关的信息，并与令牌中的声明进行比对。

总结：
服务器解析 JWT Token 的原理是验证令牌的签名，确保其完整性和真实性。然后，对令牌进行解码，获取头部和载荷的原始数据。最后，服务器验证令牌中的声明是否有效，以确定用户身份和权限等信息。JWT 的设计使得服务器在无需维护会话状态的情况下，能够验证和获取用户信息，适用于分布式、无状态的应用程序。
*/
