package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"practise-code/global"
	routerUser "practise-code/router/user"
)

func InitRouter()  *http.Server{
	Router := gin.Default()
	router := Router.Group("")
	routerUser.InitUserRouter(router)
	return &http.Server{
		Addr: ":"+global.CONFIG.Port,
		Handler: Router,
	}
}
