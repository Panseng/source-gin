package main

import (
	"fmt"
	"practise-code/global"
	"practise-code/router"
)

func main()  {
	global.InitGlobal()
	s := router.InitRouter() // 初始化路由
	err := s.ListenAndServe()
	if err != nil{
		fmt.Println(err)
	}
}