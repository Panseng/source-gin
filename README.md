# About
这是一个`gin`的实践+源码讲解项目

## 文件目录
```go
├── notes // 笔记readme
├── practise-code // 代码实践 
├── source-code  // gin源码
```

## 使用 practise-code
```bash
# 获取项目代码
git clone https://github.com/Panseng/source-gin

# 进入`practise-code`文件夹
cd source-gin/practise-code

# 依赖包安装
go mod tidy

# 编辑配置文件 ---》 config/cfg.ini
# 配置 mysql、redis、jwt等内容

# 运行项目源代码
go run main.go

# 编译
go build -o server.exe main.go

# 运行二进制
server.exe

# 打开 swagger
# http://127.0.0.1/swagger/index.html
# 注册、登录、获取用户信息
```
`practise-code`详细内容见：[readme](practise-code/readme.md)

## 源码
- [gin源码](notes/gin.md)