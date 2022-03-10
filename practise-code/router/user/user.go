package user

import (
	"github.com/gin-gonic/gin"
	"practise-code/api/v1/user"
	"practise-code/middleware"
)

func InitUserRouter(Router *gin.RouterGroup)  {
	userRouter :=Router.Group("user")
	userRouterWithAuth := Router.Group("user").Use(middleware.JWTAuth())
	userApi := user.UserApi{}
	{
		userRouter.POST("login", userApi.Login)
	}
	{
		userRouterWithAuth.GET("userinfo", userApi.GetUserInfo)
	}
}
