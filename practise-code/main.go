package main

import (
	"fmt"

	_ "practise-code/docs" // swagger
	"practise-code/global"
	"practise-code/router"
)

// @title Swagger Example API
// @version 1.0
// @description 这是一个gin实践小项目.
// @termsOfService https://github.com/Panseng
// @contact.name Jimmy
// @contact.url panseng.dr@qq.com
// @contact.email panseng.dr@qq.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:80
// @BasePath /api/v1

func main()  {
	global.InitGlobal()
	s := router.InitRouter() // 初始化路由
	err := s.ListenAndServe()
	if err != nil{
		fmt.Println(err)
	}
}