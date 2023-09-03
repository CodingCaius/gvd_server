package main

import (
	"fmt"
	"gvd_server/utils/jwts"
)

func main() {
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		NickName: "caius",
		RoleID:   2,
		UserID:   1,
	})
	fmt.Println(token, err)
}
