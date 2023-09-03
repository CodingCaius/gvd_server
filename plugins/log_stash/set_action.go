package log_stash

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gvd_server/global"
	"io"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

type Action struct {
	ip          string
	addr        string
	userName    string
	userID      uint
	serviceName string
	level       Level
	title       string
	itemList    []string
	model       *LogModel //创建之后赋值给它，用于后期更新
	token       string
	logType     LogType
}

// 设置Action对象，并初始化一些字段
func NewAction(c *gin.Context) Action {
	ip := c.RemoteIP()
	addr := "局域网"
	action := Action{
		ip:      ip,
		addr:    addr,
		logType: ActionType,
	}
	/* token := c.Request.Header.Get("token")
	jwyPayLoad := parseToken(token)
	if jwyPayLoad != nil {
		action.userID = jwyPayLoad.UserID
		action.userName = jwyPayLoad.UserName
	} */

	//不在这里解析token，现在这里拿到token
	token := c.Request.Header.Get("token")
	action.SetToken(token)
	return action
}

// 设置操作日志的级别
func (action *Action) Info(title string) {
	action.level = Info
	action.title = title
	action.save()
}
func (action *Action) Warn(title string) {
	action.level = Warning
	action.title = title
	action.save()
}
func (action *Action) Error(title string) {
	action.level = Error
	action.title = title
	action.save()
}

// 为每一个item设置具体的日志级别
func (action *Action) SetItemInfo(label string, value any) {
	action.setItem(label, value, Info)
}
func (action *Action) SetItemWarn(label string, value any) {
	action.setItem(label, value, Warning)
}
func (action *Action) SetItemErr(label string, value any) {
	action.setItem(label, value, Error)
}

// SetItem 设置一组详情
func (action *Action) SetItem(label string, value any) {
	action.setItem(label, value, Info)
}

func (action *Action) setItem(label string, value any, level Level) {
	// 判断类型
	_type := reflect.TypeOf(value).Kind()
	switch _type {
	case reflect.Struct, reflect.Map, reflect.Slice:
		// 可以设置关键字，然后有关键字的高亮显示，或者有颜色的字符
		// 颜色有两种方案
		// 1. html字符  <span style="color:red" />
		// 2. 控制字符 \033[31m xxx \033[0m
		byteData, _ := json.Marshal(value)
		action.itemList = append(action.itemList, fmt.Sprintf("<div class=\"log_item %s\"><div class=\"log_item_label\">%s</div><div class=\"log_item_content\">%s</div></div>", level.String(), label, string(byteData)))
	//case reflect.Array:
	default:
		action.itemList = append(action.itemList, fmt.Sprintf("<div class=\"log_item %s\"><div class=\"log_item_label\">%s</div><div class=\"log_item_content\">%v</div></div>", level.String(), label, value))
	}
}

// // SetItem 方法用于向 Action 结构体的 itemList 字段中添加附加项
// func (action *Action) SetItem(label string, value any) {
// 	//判断类型
// 	_type := reflect.TypeOf(value).Kind()
// 	switch _type {
// 	case reflect.Struct, reflect.Map, reflect.Slice:
// 		byteData, _ := json.Marshal(value)
// 		action.itemList = append(action.itemList, fmt.Sprintf("%s: %s", label, string(byteData)))
// 	default:
// 		action.itemList = append(action.itemList, fmt.Sprintf("%s: %v", label, value))
// 	}
// }

func (action *Action) SetToken(token string) {
	action.token = token
}

// 设置图片的访问链接
func (action *Action) SetImage(url string) {
	action.itemList = append(action.itemList, fmt.Sprintf("<div class=\"log_image\"/><img src=\"%s\"></div>", url))
}

// SetUrl 设置一组url
func (action *Action) SetUrl(title, url string) {
	// 如果要使用html显示，一定要注意xss问题
	action.itemList = append(action.itemList, fmt.Sprintf("<div class=\"log_link\"><a target=\"_blank\" href=\"%s\">%s</a></div>", url, title))
}

// SetUpload 为文件上传添加日志
func (action *Action) SetUpload(c *gin.Context) {
	// 获取表单参数
	forms, err := c.MultipartForm()
	if err != nil {
		// 设置错误信息
		action.SetItem("form参数错误", err.Error())
		return
	}
	// 遍历表单参数中的文件
	for s, headers := range forms.File {
		// 拼接信息
		action.itemList = append(action.itemList, fmt.Sprintf(
			`<div class="log_upload">
        <div class="log_upload_head">
            <span class="log_upload_file_key">%s</span>
            <span class="log_upload_file_name">%s</span>
            <span class="log_upload__file_size">%s</span>
        </div>
    </div>`, s, headers[0].Filename, FormatBytes(headers[0].Size)))
	}

}

// SetRequestHeader 设置请求头
func (action *Action) SetRequestHeader(c *gin.Context) {
	// 克隆请求头
	header := c.Request.Header.Clone()
	// 将头部数据转换为字节数组
	byteData, _ := json.Marshal(header)
	// 将字节数组追加到action.itemList中
	action.itemList = append(action.itemList, fmt.Sprintf(
		`<div class="log_request_header">
	<div class="log_request_body">
		<pre class="log_json_body">%s</pre>
	</div>
</div>`, string(byteData)))
}

// SetRequest 设置一组入参
func (action *Action) SetRequest(c *gin.Context) {
	// 请求头
	// 请求体
	// 请求路径，请求方法
	// 关于请求体的问题，拿了之后要还回去
	// 请求体读完之后就没了，为了日志和Login参数绑定时都能用，因此要还回去
	// 一定要在参数绑定之前调用
	method := c.Request.Method
	path := c.Request.URL.String()
	//对请求体进行 读取和重置
	byteData, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(byteData))

	action.itemList = append(action.itemList, fmt.Sprintf(
		`<div class="log_request">
  <div class="log_request_head">
    <span class="log_request_method">%s</span>
    <span class="log_request_path">%s</span>
  </div>
  <div class="log_request_body">
    <pre class="log_json_body">%s</pre>
  </div>
</div>`, method, path, string(byteData)))
}

// SetResponse 设置一组出参
func (action *Action) SetResponse(c *gin.Context) {
	c.Set("action", action)
}

func (action *Action) SetResponseContent(response string) {
	action.itemList = append(action.itemList, fmt.Sprintf(`
<div class="log_response">
	<pre class="log_request_json">%s</pre>
</div>
`, response))

	//<div class="log_response">
	//<pre class="log_request_json">{"code":0,"data":"/uploads/caius/4b2185c792a7c409b0bf6f860fdd0bef.jpg","msg":"图片上传成功"}</pre>
	//</div>
}

func (action *Action) SetFlush() {
	action.level = action.model.Level
	action.save()
}

// 用来保存日志记录到数据库中
func (action *Action) save() {
	content := strings.Join(action.itemList, "\n")

	//有些接口是不需要token的，比如游客访问文档
	if action.token != "" {
		//留到这里解析token，顺便拿到用户id和用户名
		jwyPayLoad := parseToken(action.token)
		if jwyPayLoad != nil {
			action.userID = jwyPayLoad.UserID
			action.userName = jwyPayLoad.UserName
		}
	}

	//这一步，model为空的话就创建一个并赋值，用于之后判断
	if action.model == nil {
		action.model = &LogModel{
			IP:          action.ip,
			Addr:        action.addr,
			Level:       action.level,
			Title:       action.title,
			Content:     content, //第一次的content
			UserID:      action.userID,
			UserName:    action.userName,
			ServiceName: action.serviceName,
			//这里不能写死
			Type: action.logType,
		}
		global.DB.Create(action.model)
		// 如果不对content进行置空，那么content会重复
		action.itemList = []string{}
		return
	}
	//如果model不为空，说明是同一个log，就执行更新操作
	global.DB.Model(action.model).Updates(LogModel{
		Level: action.level,
		Title: action.title,
		//原来的content 加上 新的 content
		Content: action.model.Content + "\n" + content,
	})
}
