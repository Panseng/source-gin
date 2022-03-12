package user

import (
	"fmt"
	"practise-code/global"
	"practise-code/model/types"
	"practise-code/utils"
)

func Login(u types.User) ( userInfo types.User, err error) {
	if global.SQLHARD.Db == nil{
		return types.User{}, fmt.Errorf("database do not init")
	}

	// 这里假定前端加密的密码已经解密（如果有加密）
	u.Password = utils.MD5(u.Password)
	err = global.SQLHARD.Db.Where("username = ? AND password = ?", u.Username, u.Password).First(&userInfo).Error
	return userInfo, err
}
