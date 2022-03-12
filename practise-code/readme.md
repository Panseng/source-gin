# About
gin框架的实践项目

## 目录
```go
├── api
│   └── v1
├── config  // 配置包
├── docs  // swagger api生成文件
├── global  // 全局对象
├── middleware // 中间件
├── model  // 模型层
│   ├── http // http 入参、出参结构体
│   └── sql  // sql 数据结构体 
├── router // 路由层
├── static // 静态文件
└── utils
    ├── response // 回复消息
    └── validator // 值校验
```

## 技术栈
主体框架：`gin-gonic/gin`

配置项读取：`gopkg.in/ini.v1`

sql数据：`gorm.io/gorm`（`jinzhu/gorm`的更新版本）

用户认证（中间件）：`golang-jwt/jwt/v4`

API文档 & API调测：`swaggo/swag`

日志：`go.uber.org/zap`

缓存：`go-redis/redis/v8`

## swagger 尾部参数
如以下路径的id，需要作为api值传递
```go
Router.GET("userinfo/:id", userApi.GetUserInfo)
```
则需要在handler中，参数项类型设置为path，同时在路径加上变量名称`{id}`
```go
// @Param id path uint true "用户id"
// @Router /user/userinfo/{id} [get]
```