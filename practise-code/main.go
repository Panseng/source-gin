package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"practise-code/config"
	"practise-code/global"
	routerUser "practise-code/router/user"
)

func main()  {
	global.CONFIG = config.Config{
		JWT: config.JWT{
			SigningKey: "DSdf1s15afSDA5sd1fdf5sdf1f",
			ExpiresTime: 1000*60*10, // 10分钟
			Issuer: "JP",
		},
	}
	Router := gin.Default()
	router := Router.Group("")
	routerUser.InitUserRouter(router)
	s := &http.Server{
		Addr: ":8080",
		Handler: Router,
	}
	err := s.ListenAndServe()
	if err != nil{
		fmt.Println(err)
	}
}