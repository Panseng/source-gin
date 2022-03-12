package middleware

import (
	"github.com/gin-gonic/gin"
	"practise-code/global"

	"practise-code/utils"
	"practise-code/utils/response"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		if token == ""{
			global.SLOG.Infow("JWT权限校验，未登录或非法访问，jwt为空")
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		claims, err := utils.ParseToken(token)
		if err != nil{
			if err == utils.TokenExpired {
				global.SLOG.Infow("JWT权限校验，授权已过期")
				response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
				c.Abort()
				return
			}
			global.SLOG.Warnw("JWT权限校验，其他错误，token: " + token + ", err: " + err.Error())
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}
		cacheJWT := ""
		// 不支持多点登录，则获取缓存jwt
		if !global.CONFIG.UseMultipoint{
			cacheJWT, err = global.REDIS.GetRedisJWT(claims.Username)
			if err != nil{
				global.SLOG.Errorf("redis get jwt err: %+v", err)
			}
		}
		if cacheJWT != "" && cacheJWT != token{
			global.SLOG.Infow("JWT权限校验，当前授权已作废")
			response.FailWithDetailed(gin.H{"reload": true}, "用户已在其他窗口登录，当前授权已作废", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
