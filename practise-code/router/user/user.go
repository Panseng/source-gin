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
	{ // 不需要用户权限的页面
		userRouter.POST("register", userApi.Register)
		userRouter.POST("login", userApi.Login)
	}
	{ // 需要用户权限的页面
		userRouterWithAuth.GET("userinfo/:uid", userApi.GetUserInfo)
	}
}
