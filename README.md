# gvd_server
[网站地址](http://docs.codingcaius.top/)  
本项目是使用 Golang 语言开发的基于 Gin 实现的前后端分离的文档管理系统，集成了 JWT 鉴权，文档和用户的添加修改和删除，权限管理，日志管理，图片管理，路由管理，redis 缓存，全文搜索，定时任务等功能。 
## 使用的技术

- Golang 
- MySQL 
- Redis 
- Gin
- Gorm
- Jwt  
- Swagger
- logrus 
- ElasticSearch
- Docker

## 主要功能

- **数据库管理**：使用MySQL进行数据存储，利用Gorm进行ORM操作，实现文档和用户的添加、修改和删除等功能。
- **权限管理**：基于JWT实现用户身份验证和授权，通过角色进行用户权限管理，确保不同角色用户拥有不同的操作权限。
- **加密和预览**：为文档添加加密功能，支持为不同角色的用户设置不同的文档密码和预览权限。
- **日志管理**：记录登录日志、操作日志和运行日志，利用logrus进行日志记录与管理。
- **图片管理**：支持上传和删除图片，限制图片格式和大小，确保系统安全性和稳定性。
- **缓存管理**：利用Redis实现对Token、文档和浏览量等数据的缓存，提高系统性能和响应速度。
- **中间件管理**：实现用户登录中间件、超级管理员鉴权中间件和日志中间件，提升系统安全性和可维护性。
- **全文搜索**：通过ElasticSearch实现全文搜索功能，支持用户快速查找和定位文档内容。
- **定时任务**：每晚两点自动同步文档的浏览和点赞数据，保证数据的准确性和完整性



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
├── plugins          插件，里面是独立的日志系统
├── routers          路由
├── service          服务
├── setting.yaml     配置文件
├── testdata         测试用例
├── uploads          上传的文件
└── utils            一些工具
  
</pre></code>
</details>



## 表结构
![t](https://github.com/CodingCaius/gvd_server/blob/master/uploads/20230803113621.png)
<details>
<summary>展开查看</summary>


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

## 中间件

用户登录中间件





超级管理员鉴权中间件





日志中间件



## 解耦的日志系统


## 项目截图
<details>
<summary>展开查看</summary>

![t](https://github.com/CodingCaius/gvd_server/blob/master/uploads/QQ%E6%88%AA%E5%9B%BE20240304105637.png)  

![t](https://github.com/CodingCaius/gvd_server/blob/master/uploads/QQ%E6%88%AA%E5%9B%BE20240304105833.png)  

### 角色管理
![t](https://github.com/CodingCaius/gvd_server/blob/master/uploads/QQ%E6%88%AA%E5%9B%BE20240304105934.png)

### 日志管理
登录日志
![t](https://github.com/CodingCaius/gvd_server/blob/master/uploads/%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE%202024-02-10%20145145.png)
操作日志
![t](https://github.com/CodingCaius/gvd_server/blob/master/uploads/%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE%202024-02-10%20145123.png)
![t](https://github.com/CodingCaius/gvd_server/blob/master/uploads/%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE%202024-02-10%20145240.png)

### 图片管理
![t](https://github.com/CodingCaius/gvd_server/blob/master/uploads/QQ%E6%88%AA%E5%9B%BE20240304111233.png)

</details>


## 部署和运行

如果是 windows 下开发，执行一下命令进行交叉编译

```bash
SET CGO_ENABLED=0
set GOARCH=amd64
set GOOS=linux
go build -o main
set GOOS=windows
```

linux 环境下直接 `go build -o main` 

然后将编译好的 main 文件和配置文件 settings.yaml 放入服务器对应目录

比如：

```bash
opt
  gin-vue-docs
    server
      main // 后端打包之后的文件，放上服务器上记得执行 chmod +x main
      settings.yaml
      uploads  // 后端的一些用户上传文件目录
      gvd_db_20230826.sql
      gvd_server_full_text_index_20230826.json
    web
      dist // 前端打包之后的目录
    gvd_server.ini
    nginx.conf

```

先cd到/opt/gin-vue-docs/server

```bash
cd /opt/gin-vue-docs/server
```



### docker配置mysql

```bash
docker pull mysql:5.7

mkdir -p ./mysql/conf ./mysql/datadir

docker run -itd --name mysql --restart=always -p 3306:3306 -v /opt/gin-vue-docs/server/mysql/conf:/etc/mysql/conf.d -v /opt/gin-vue-docs/server/mysql/datadir:/var/lib/mysql -e  MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=gvd_db mysql:5.7
```

创建一个gvd_db 的数据库，root密码是root，端口映射为本机的3306映射到容器的3306

如果已经有mysql服务了，那换一个端口就行了，在对应的settings.yaml配置文件修改就ok了

### docker配置redis

```bash
docker pull redis:5.0.5

docker run -itd --name redis --restart=always -p 6379:6379 redis:5.0.5 --requirepass "redis"
```



### docker配置es

```bash
mkdir -p es/config & mkdir -p es/data & mkdir -p es/plugins
chmod 777 es/data
echo "http.host: 0.0.0.0" > es/config/elasticsearch.yml

docker run --name docs_es -p 9200:9200   -e "discovery.type=single-node" -e ES_JAVA_OPTS="-Xms84m -Xmx512m" -v /opt/gin-vue-docs/server/es/config/elasticsearch.yml:/usr/share/elasticsearch/config/elasticsearch.yml -v /opt/gin-vue-docs/server/es/data:/usr/share/elasticsearch/data -v /opt/gin-vue-docs/server/es/plugins:/usr/share/elasticsearch/plugins -d elasticsearch:7.12.0

```



### 修改 settings文件

根据具体情况修改 settings文件，比如：

```yaml
system:
    ip: 127.0.0.1
    port: 8082
    env: dev
mysql:
    host: 127.0.0.1
    port: 3306
    config: charset=utf8mb4&parseTime=True&loc=Local
    db: gvd_db
    username: root
    password: root
    logLevel: error
redis:
    ip: 127.0.0.1
    port: 6379
    password: "redis"
    poolSize: 100
es:
    addr: http://127.0.0.1:9200
    user: 
    password: 
jwt:
    expires: 8
    issuer: caius
    secret: soagfeohfscz
site:
    title: caius docs
    icon: 
    abstract: 
    iconHref:
    href: ""
    footer: ""
    content: ""

```



### 运行

```
./main
```








