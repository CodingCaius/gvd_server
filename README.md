# gvd_server
gin-vue-doc 前后端分离的文档系统  
[网站地址](http://docs.codingcaius.top/)



## 项目目录
<details>
<summary>展开查看</summary>
<pre><code>
├── api             通过api接口调用方法
├── config          配置文件中映射的结构体
├── core            初始化连接的一些操作
├── docs            swagger api文档
├── flags           命令行参数绑定
├── global          全局变量
├── go.mod        
├── go.sum
├── logs             日志文件
├── main.go          主函数
├── middleware       gin 的中间件
├── models           表结构
├── plugins          插件
├── routers          路由
├── service          服务
├── setting.yaml     配置文件
├── testdata         测试用例
├── uploads          上传的文件
└── utils            一些工具
  
</pre></code>
</details>



## 表结构
<details>
<summary>展开查看</summary>
![t](https://github.com/CodingCaius/gvd_server/blob/master/uploads/20230803113621.png)

### 角色表

```Go
package models

type RoleModel struct {
  Model
  Title    string     `gorm:"size:16;not null;comment:角色名称" json:"title"`                                    // 角色的名称
  Pwd      string     `gorm:"size:64;comment:角色的密码" json:"-"`                                                // 角色密码
  IsSystem bool       `gorm:"column:isSystem;comment:是否是系统角色" json:"isSystem"`                               // 是否是系统角色
  DocsList []DocModel `gorm:"many2many:role_doc_models;joinForeignKey:RoleID;JoinReferences:DocID" json:"-"` // 角色拥有的文档列表
}

```



### 文档表

```Go
package models

type DocModel struct {
  Model
  Title           string      `gorm:"comment:文档标题" json:"title"`
  Content         string      `gorm:"comment:文档内容" json:"-"`
  DiggCount       int         `gorm:"comment:点赞量;column:diggCount" json:"diggCount"`
  LookCount       int         `gorm:"comment:浏览量;column:lookCount" json:"lookCount"`
  Key             string      `gorm:"comment:key;not null;unique" json:"key"`
  ParentID        *uint       `gorm:"comment:父文档id;column:parentID" json:"parentID"`
  ParentModel     *DocModel   `gorm:"foreignKey:ParentID" json:"-"` // 父文档
  Child           []*DocModel `gorm:"foreignKey:ParentID" json:"-"` // 它会有子孙文档
  FreeContent     string      `gorm:"comment:预览部分;column:freeContent" json:"freeContent"`
  UserCollDocList []UserModel `gorm:"many2many:user_coll_doc_models;joinForeignKey:DocID;JoinReferences:UserID" json:"-"`
}

```

### 角色文档表

```Go
package models

type RoleDocModel struct {
  Model
  RoleID      uint      `gorm:"column:roleID;comment:角色id" json:"roleID"`
  RoleModel   RoleModel `gorm:"foreignKey:RoleID" json:"-"`
  DocID       uint      `gorm:"column:docID;comment:文档id" json:"docID"`
  DocModel    DocModel  `gorm:"foreignKey:DocID" json:"-"`
  Pwd         *string   `gorm:"column:pwd;comment:密码配置" json:"pwd"`                 // null ""  "有值"  优先级： 角色文档密码 > 角色密码
  FreeContent *string   `gorm:"column:freeContent;comment:试看配置" json:"freeContent"` // 试看部分 优先级：角色文档试看  > 文档试看字段 > 文档按照特殊字符分隔的试看
  Sort        int       `gorm:"column:sort;comment:排序" json:"sort"`                 // 排序
}

```



### 用户表

```Go
package models

type UserModel struct {
  Model
  UserName  string    `gorm:"column:userName;size:36;unique;not null;comment:用户名" json:"-"` // 用户名
  Password  string    `gorm:"column:password;size:128;comment:密码"  json:"-"`                // 密码
  Avatar    string    `gorm:"column:avatar;size:256;comment:头像"  json:"avatar"`             // 头像
  NickName  string    `gorm:"column:nickName;size:36;comment:昵称"  json:"nickName"`          // 昵称
  Email     string    `gorm:"column:email;size:128;comment:邮箱"  json:"email"`               // 邮箱
  Token     string    `gorm:"column:token;size:64;comment:其他平台的唯一id"  json:"-"`             // 其他平台的唯一id
  IP        string    `gorm:"column:ip;size:16;comment:ip地址"  json:"ip"`                    // ip
  Addr      string    `gorm:"column:addr;size:64;comment:地址"  json:"addr"`                  // 地址
  RoleID    uint      `gorm:"column:roleID;comment:用户对应的角色" json:"roleID"`                  // 用户对应的角色
  RoleModel RoleModel `gorm:"foreignKey:RoleID" json:"-"`
}

```

### 用户收藏文档表

```Go
package models

type UserCollDocModel struct {
  Model
  DocID     uint      `gorm:"column:docID" json:"docID"`
  DocModel  DocModel  `gorm:"foreignKey:DocID"`
  UserID    uint      `gorm:"column:userID" json:"userID"`
  UserModel UserModel `gorm:"foreignKey:UserID"`
}

```



### 用户密码访问文档表

```Go
package models

type UserPwdDocModel struct {
  Model
  UserID uint `gorm:"column:userID" json:"userID"`
  DocID  uint `gorm:"column:docID" json:"docID"`
}

```

### 图像表

```Go
package models

import "fmt"

type ImageModel struct {
  Model
  UserID    uint      `gorm:"column:userID;comment:用户id" json:"userID"`
  UserModel UserModel `gorm:"foreignKey:UserID" json:"-"`
  FileName  string    `gorm:"column:fileName;size:64；comment:文件名" json:"fileName"`
  Size      int64     `gorm:"column:size;comment:文件大小，单位字节" json:"size"`
  Path      string    `gorm:"column:path;size:128;comment:文件路径" json:"path"`
  Hash      string    `gorm:"column:hash;size:64;comment:文件的hash" json:"hash"`
}

func (image ImageModel) WebPath() string {
  return fmt.Sprintf("/%s", image.Path)
}

```

### 登录记录表

```Go
package models

// LoginModel 用户登录数据
type LoginModel struct {
  Model
  UserID    uint      `gorm:"column:userID" json:"userID"`
  UserModel UserModel `gorm:"foreignKey:UserID" json:"-"`
  IP        string    `gorm:"size:20" json:"ip"` // 登录的ip
  NickName  string    `gorm:"column:nickName;size:42" json:"nickName"`
  UA        string    `gorm:"size:256" json:"ua"` // ua
  Token     string    `gorm:"size:256" json:"token"`
  Device    string    `gorm:"size:256" json:"device"` // 登录设备
  Addr      string    `gorm:"size:64" json:"addr"`
}

```

### 文档数据表

```Go
package models

// DocDataModel 文档数据表
type DocDataModel struct {
  Model
  DocID     uint   `gorm:"column:docID" json:"docID"`
  DocTitle  string `gorm:"column:docTitle" json:"docTitle"`
  LookCount int    `gorm:"column:lookCount" json:"lookCount"`
  DiggCount int    `gorm:"column:diggCount" json:"diggCount"`
  CollCount int    `gorm:"column:collCount" json:"collCount"`
}

```


</details>
