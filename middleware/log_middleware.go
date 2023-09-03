
//用于在 处理 HTTP 请求和响应过程中，记录请求和响应的内容，然后将这些内容传递给一个日志记录插件 log_stash



package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"gvd_server/plugins/log_stash"
)

type responseWrite struct {
	gin.ResponseWriter
	byteData *bytes.Buffer
}

func (rw responseWrite) Write(buf []byte) (int, error) {
	rw.byteData.Write(buf)
	return rw.ResponseWriter.Write(buf)
}

func LogMiddleWare() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		// 请求
		r := responseWrite{
			ResponseWriter: c.Writer,
			byteData:       bytes.NewBuffer([]byte{}),
		}
		c.Writer = r
		c.Next()
		// 响应
		_action, ok := c.Get("action")
		if !ok {
			return
		}
		action, ok := _action.(*log_stash.Action)
		if !ok {
			return
		}

		action.SetResponseContent(r.byteData.String())
		action.SetFlush()
	}
}