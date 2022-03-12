package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"practise-code/model/types"
	"strconv"
	"time"

	"practise-code/global"
	httpUser "practise-code/model/http/user"
	sqlUser "practise-code/model/sql/user"
	"practise-code/utils"
	utilsResponse "practise-code/utils/response"
	utilsValid "practise-code/utils/validator"
)

type UserApi struct{}

// @Summary 用户注册
// @Description 用户注册，向服务器提供用户数据
// @Tags user
// @Accept json
// @Param data body httpUser.RegisterRequest true "用户名, 密码"
// @Success 200 {string} json "注册成功信息"
// @Failure 200 {string} json "错误提示信息"
// @Router /user/register [post]
func (u *UserApi) Register(c *gin.Context){
	var r httpUser.RegisterRequest
	_ = c.ShouldBindJSON(&r)
	if err := utilsValid.Verify(r, utilsValid.RegisterVerify); err != nil{
		global.SLOG.Infof("用户注册，用户数据不完整：user-%+v, err-%+v", r, err)
		utilsResponse.FailWithMessage(err.Error(), c)
		return
	}

	// 这里在思考，将sql层单独放到model是否合适
	if sqlUser.CheckUserName(r.Username) != gorm.ErrRecordNotFound{
		global.SLOG.Infof("用户注册，用户名不可用：user-%+v", r)
		utilsResponse.FailWithMessage("用户名不可用", c)
		return
	}
	user := types.User{
		Username: r.Username,
		Password: r.Password,
	}
	if err := sqlUser.Register(user); err == nil{
		utilsResponse.OkWithMessage("成功注册", c)
		return
	} else {
		global.SLOG.Warnf("用户注册，用户注册失败：user-%+v, err-%+v", r, err)
		utilsResponse.FailWithMessage(err.Error(), c)
		return
	}
}

// @Summary 用户登录
// @Description 用户登录，向服务器提供用户登录数据
// @Tags user
// @Accept json
// @Param data body httpUser.LoginRequest true "用户名, 密码, 验证码, 验证码ID"
// @Success 200 {object} response.Response{data=httpUser.LoginResponse,msg=string} "用户注册账号,返回包括用户信息"
// @Failure 200 {string} json "错误提示信息"
// @Router /user/login [post]
func (u *UserApi) Login(c *gin.Context) {
	var l httpUser.LoginRequest
	_ = c.ShouldBindJSON(&l)                                             // 获取值
	if err := utilsValid.Verify(l, utilsValid.LoginVerify); err != nil { // 校验取值的规范性
		global.SLOG.Infof("登录，用户数据不完整：user-%+v, err-%+v", l, err)
		utilsResponse.FailWithMessage(err.Error(), c)
		return
	}
	var user types.User
	var err error
	// 从数据库校验值对错
	if user, err = sqlUser.Login(types.User{Username: l.Username, Password: l.Password}); err != nil{
		if err == gorm.ErrRecordNotFound{
			global.SLOG.Infof("登录，用户密码错误或不存在：user-%+v, err-%+v", l, err)
			utilsResponse.FailWithMessage("用户密码错误或不存在", c)
			return
		}
		global.SLOG.Warnf("登录，其他错误：user-%+v, err-%+v", l, err)
		utilsResponse.FailWithMessage("数据库未初始化", c)
		return
	}

	expiresTime := global.CONFIG.JWT.ExpiresTime
	accessToken, err := utils.GetToken(user)
	if err != nil {
		global.SLOG.Error("登录，其他错误：user-%+v, err-%+v", user, err)
		utilsResponse.FailWithMessage("获取token失败", c)
		return
	}
	// 不允许用户多点登录的情况下，缓存最新的jwt
	// 配合websocket，可以即时将其他点用户登出
	if !global.CONFIG.UseMultipoint{
		err = global.REDIS.SetRedisJWT(accessToken, user.Username, time.Duration(expiresTime)*time.Second)
		if err != nil{
			global.SLOG.Errorf("redis cache jwt err: %+v", err)
		}
	}
	utilsResponse.OkWithDetailed(httpUser.LoginResponse{
		Token:     accessToken,
		ExpiresAt: expiresTime,
	}, "登录成功", c)
}

// @Summary 获取用户信息
// @Description 获取用户信息
// @Tags user
// @Accept json
// @Success 200 {object} response.Response{data=httpUser.UserInfoResponse,msg=string} "用户信息"
// @Failure 200 {string} json "错误提示信息"
// @Param x-token header string true "jwt"
// @Param id path uint true "用户id"
// @Router /user/userinfo/{id} [get]
func (u *UserApi) GetUserInfo(c *gin.Context) {
	uid, err := strconv.ParseUint(c.Param("uid"), 10, 64)
	if err != nil{
		global.SLOG.Infof("获取用户信息，从路径获取用户id出错：%+v", err)
	}
	user, err := sqlUser.GetUserInfo(uint(uid))
	if err != nil{
		if err == gorm.ErrRecordNotFound{
			global.SLOG.Infof("获取用户信息，用户不存在：uid-%v, err-%+v", uid, err)
			utilsResponse.FailWithMessage("用户不存在", c)
			return
		}
		global.SLOG.Warnf("获取用户信息，其他错误：uid-%v, err-%+v", uid, err)
		utilsResponse.FailWithMessage(err.Error(), c)
		return
	}
	utilsResponse.OkWithDetailed(gin.H{"userInfo": httpUser.UserInfoResponse{
		ID: user.ID,
		Name: user.Username,
		Mobile: user.Phone,
		HeaderImg: user.HeaderImg,
	}}, "获取成功", c)
}
