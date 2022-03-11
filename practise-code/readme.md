# About
gin框架的实践项目

## 目录
```go
├── api
│   └── v1
├── config  // 配置包
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

## 功能简介
主体框架：`gin-gonic/gin`

配置项读取：`gopkg.in/ini.v1`

sql数据：`gorm.io/gorm`（`jinzhu/gorm`的更新版本）

用户认证（中间件）：`golang-jwt/jwt/v4`

## Todo
- [ ] swagger 生成api文档
- [ ] redis 缓存 jwt