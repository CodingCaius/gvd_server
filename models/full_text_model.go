package models

type ESIndexInterFace interface {
	Index() string
	Mapping() string
}

// FullTextModel es全文搜索的model
type FullTextModel struct {
	DocID uint   `json:"docID"`
	ID    string `json:"id"`    // es的id
	Title string `json:"title"` // 标题
	Body  string `json:"body"`  // 正文
	Slug  string `json:"slug"`  // 跳转地址，可以由docID和title拼接而来
}

func (FullTextModel) Index() string {
	return "gvd_server_full_text_index"
}

// 在 "properties" 下定义了四个字段：
// "body"：这是一个名为 "body" 的字段，类型为 "text"。这通常用于存储文本内容。
// "title"：这是一个名为 "title" 的字段，类型为 "text"，同时也定义了一个 "keyword" 子字段，类型为 "keyword"，并设置了 "ignore_above" 参数为 256。这样的设置通常用于存储标题，并且 "keyword" 子字段用于精确匹配和聚合，而 "text" 字段用于全文搜索。
// "slug"：这是一个名为 "slug" 的字段，类型为 "keyword"。通常用于存储用于标识文档的唯一标识符。
// "docID"：这是一个名为 "docID" 的字段，类型为 "integer"，通常用于存储整数类型的标识符或ID。

// 定义了一个 Elasticsearch 索引的映射，其中包括了用于存储文本内容、标题、唯一标识符和整数标识符的字段
func (FullTextModel) Mapping() string {
	return `
{
  "mappings": {
    "properties": {
      "body": {
        "type": "text"
      },
      "title": {
        "type": "text",
	      "fields": {
          "keyword": {
              "type": "keyword",
			        "ignore_above": 256
          }
        }
	    },
	    "slug": {
        "type": "keyword"
      },
      "docID": {
        "type": "integer"
      }
    }
  }
}
`
}
