package full_search_service

import (
	"fmt"
	"strings"
)

type SearchData struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Slug  string `json:"slug"`
}

// MarkdownParse 解析markdown的正文，传递文档标题，和正文
func MarkdownParse(id uint, title string, content string) (searchDataList []SearchData) {
	var body string
	var headList, bodyList []string
	var isCode bool
	list := strings.Split(content, "\n")
	headList = append(headList, title)
	for _, s := range list {
		if strings.HasPrefix(s, "```") {
			isCode = !isCode
		}
		if strings.HasPrefix(s, "#") && !isCode {
			// 标题
			headList = append(headList, getHead(s))
			bodyList = append(bodyList, body)
			body = ""
			continue
		}
		body += s
	}
	bodyList = append(bodyList, body)
	ln := len(headList)
	for i := 0; i < ln; i++ {
		searchDataList = append(searchDataList, SearchData{
			Title: headList[i],
			Body:  bodyList[i],
			Slug:  getSlug(id, headList[i]),
		})
	}
	return
}

// getSlug 拼接slug
func getSlug(id uint, title string) string {
	// 中间的#号是前端的锚点
	return fmt.Sprintf("%d#%s", id, title)
}

// getHead 处理head
func getHead(head string) string {
	head = strings.ReplaceAll(head, "#", "")
	head = strings.ReplaceAll(head, " ", "")
	return head
}