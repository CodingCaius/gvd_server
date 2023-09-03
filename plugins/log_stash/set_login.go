package log_stash

import (
	"gvd_server/global"

	"github.com/gin-gonic/gin"
)

// NewSuccessLogin 登录成功的日志
func NewSuccessLogin(c *gin.Context) {
	token := c.Request.Header.Get("token")
	jwyPayLoad := parseToken(token)
	saveLoginLog("登录成功", "--", jwyPayLoad.UserID, jwyPayLoad.UserName, true, c)
}

// NewFailLogin 登录失败的日志
func NewFailLogin(title, userName, pwd string, c *gin.Context) {
	//登陆失败不能拿到userid，赋0
	saveLoginLog(title, pwd, 0, userName, false, c)
}

func saveLoginLog(title string, content string, userID uint, userName string, status bool, c *gin.Context) {
	ip := c.RemoteIP()
	addr := "局域网"
	global.DB.Create(&LogModel{
		IP:       ip,
		Addr:     addr,
		Title:    title,
		Content:  content,
		UserID:   userID,
		UserName: userName,
		Status:   status,
		Type:     LoginType,
	})
}
