package user

import (
	"github.com/gin-gonic/gin"

	"practise-code/global"
	"practise-code/model"
	httpUser "practise-code/model/http/user"
	sqlUser "practise-code/model/sql/user"
	"practise-code/utils"
	utilsResponse "practise-code/utils/response"
	utilsValid "practise-code/utils/validator"
)

type UserApi struct{}

func (u *UserApi) Login(c *gin.Context) {
	var l httpUser.LoginRequest
	_ = c.ShouldBindJSON(&l)                                             // 获取值
	if err := utilsValid.Verify(l, utilsValid.LoginVerify); err != nil { // 校验取值的规范性
		utilsResponse.FailWithMessage(err.Error(), c)
		return
	}
	// Todo
	// 从数据库校验值对错

	expiresTime := global.CONFIG.JWT.ExpiresTime
	// Todo
	// 此处暂时用手写user代替
	user := sqlUser.User{
		DefaultField: model.DefaultField{
			ID: 1,
		},
		Username: "test",
	}
	accessToken, err := utils.GetToken(user)
	if err != nil {
		utilsResponse.FailWithMessage("获取token失败", c)
		return
	}
	utilsResponse.OkWithDetailed(httpUser.LoginResponse{
		Token:     accessToken,
		ExpiresAt: expiresTime,
	}, "登录成功", c)
}

func (u *UserApi) GetUserInfo(c *gin.Context) {
	// 假装有mysql
	utilsResponse.OkWithDetailed(gin.H{"userInfo": httpUser.UserInfoResponse{
		ID: 1,
		Name: "test",
		Mobile: "12546597985",
		HeaderImg: "https://qmplusimg.henrongyi.top/gva_header.jpg",
	}}, "获取成功", c)
}
