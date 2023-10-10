# gvd_server
gin-vue-doc 前后端分离的文档系统
# 目录结构


```
.
├── api
│   ├── data_api
│   │   ├── data_login_date.go
│   │   ├── data_look_date.go
│   │   ├── data_sum.go
│   │   └── entry.go
│   ├── doc_api
│   │   ├── doc_content.go
│   │   ├── doc_create.go
│   │   ├── doc_digg.go
│   │   ├── doc_edit_content.go
│   │   ├── doc_info.go
│   │   ├── doc_pwd.go
│   │   ├── doc_remove.go
│   │   ├── doc_search.go
│   │   ├── doc_update.go
│   │   └── entry.go
│   ├── entry.go
│   ├── image_api
│   │   ├── entry.go
│   │   ├── image_list.go
│   │   ├── image_remove.go
│   │   └── image_update.go
│   ├── log_api
│   │   ├── entry.go
│   │   ├── log_list.go
│   │   ├── log_read.go
│   │   └── log_remove.go
│   ├── role_api
│   │   ├── entry.go
│   │   ├── role_create.go
│   │   ├── role_id_list.go
│   │   ├── role_list.go
│   │   ├── role_remove.go
│   │   └── role_update.go
│   ├── role_doc_api
│   │   ├── entry.go
│   │   ├── role_doc_create.go
│   │   ├── role_doc_info.go
│   │   ├── role_doc_info_update.go
│   │   ├── role_doc_list.go
│   │   ├── role_doc_remove.go
│   │   ├── role_doc_tree.go
│   │   └── role_doc_update.go
│   ├── site_api
│   │   ├── entry.go
│   │   ├── site_detail.go
│   │   └── site_update.go
│   ├── user_api
│   │   ├── entry.go
│   │   ├── user_create.go
│   │   ├── user_info.go
│   │   ├── user_list.go
│   │   ├── user_login.go
│   │   ├── user_logout.go
│   │   ├── user_remove.go
│   │   ├── user_update.go
│   │   ├── user_update_info.go
│   │   └── user_update_password.go
│   └── user_center_api
│       ├── entry.go
│       ├── user_coll_doc.go
│       └── user_coll_doc_list.go
├── config
│   ├── config_es.go
│   ├── config_jwt.go
│   ├── config_mysql.go
│   ├── config_redis.go
│   ├── config_site.go
│   ├── config_system.go
│   └── enter.go
├── core
│   ├── init_addr_db.go
│   ├── init_es.go
│   ├── init_logrus.go
│   ├── init_mysql.go
│   ├── init_redis.go
│   └── setting.go
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── flags
│   ├── entry.go
│   ├── flag_es_index.go
│   ├── flags_db.go
│   ├── flags_dump.go
│   ├── flags_es_dump.go
│   ├── flags_es_load.go
│   ├── flags_load.go
│   └── flags_port.go
├── global
│   └── global.go
├── go.mod
├── go.sum
├── logs
├── main
├── main.go
├── middleware
│   ├── jwt_admin.go
│   ├── jwt_auth.go
│   └── log_middleware.go
├── models
│   ├── doc_data_model.go
│   ├── doc_model.go
│   ├── entry.go
│   ├── full_text_model.go
│   ├── image_model.go
│   ├── login_model.go
│   ├── role_doc_model.go
│   ├── role_model.go
│   ├── user_coll_doc_model.go
│   ├── user_model.go
│   └── user_pwd_doc_model.go
├── plugins
│   └── log_stash
│       ├── level.go
│       ├── log_type.go
│       ├── model.go
│       ├── parse_token.go
│       ├── set_action.go
│       ├── set_login.go
│       ├── set_runtime.go
│       └── utils.go
├── routers
│   ├── data_router.go
│   ├── doc_router.go
│   ├── enter.go
│   ├── image_router.go
│   ├── log_router.go
│   ├── role_doc_router.go
│   ├── role_router.go
│   ├── site_router.go
│   ├── user_center_router.go
│   └── user_router.go
├── service
│   ├── common
│   │   ├── list
│   │   │   └── query_list.go
│   │   └── res
│   │       └── entry.go
│   ├── cron_service
│   │   ├── entry.go
│   │   ├── sync_doc_data_date.go
│   │   └── sync_doc_data.go
│   ├── es_service
│   │   └── indexs
│   │       └── entry.go
│   ├── full_search_service
│   │   ├── full_search_create.go
│   │   ├── full_search_delete.go
│   │   ├── full_search_update.go
│   │   └── markdown_parse.go
│   └── redis_service
│       ├── redis_count.go
│       ├── redis_doc_content.go
│       ├── redis_logout.go
│       └── redis_role_doc_tree.go
├── setting.yaml
├── testdata
│   ├── 定时任务
│   │   └── main.go
│   ├── 日志样式
│   │   └── 日志样式.html
│   ├── 生成jwt.go
│   └── 通过ip地址算物理地址.go
├── uploads
└── utils
    ├── file
    │   └── format_bytes.go
    ├── hash
    │   └── md5.go
    ├── ip
    │   └── get_addr.go
    ├── jwts
    │   ├── entry.go
    │   ├── generate_token.go
    │   └── parse_token.go
    ├── pwd
    │   └── pwd.go
    ├── set
    │   ├── set_sub2.go
    │   ├── set_sub.go
    │   └── set_union.go
    ├── utils.go
    └── valid
        └── valid.go
```
