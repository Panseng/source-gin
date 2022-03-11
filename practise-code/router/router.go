package router

import (
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"

	"github.com/gin-gonic/gin"

	"practise-code/global"
	routerUser "practise-code/router/user"
)

func InitRouter()  *http.Server{
	Router := gin.Default()

	// 添加 swagger
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router := Router.Group("/api/v1")
	routerUser.InitUserRouter(router)
	return &http.Server{
		Addr: ":"+global.CONFIG.Port,
		Handler: Router,
	}
}
