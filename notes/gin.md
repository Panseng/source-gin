# Gin源码
以practise-code的路由为例：[router.go](../practise-code/router/router.go)

## gin.Default
1. 打印提示信息：版本号检测 & 预填的中间件

2. 初始化一个路由引擎
   - 如果当前是开发 / 调试环境，则会打印提示信息
   - 如果不需要 gin 的打印信息可以参考一下代码进行设置
    ```go
    // using env:	export GIN_MODE=release
    // using code:	gin.SetMode(gin.ReleaseMode)
    gin.SetMode(gin.ReleaseMode)
    gin.DefaultWriter = ioutil.Discard
    ```
3. 填入日志打印、恢复中间件

4. 返回引擎`Engine`
   - 主要包括路由构建、最终的路由树、模板解析等结构体、功能
   
## 
