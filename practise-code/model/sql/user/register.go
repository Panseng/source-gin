package user

import (
	"practise-code/global"
	"practise-code/model/types"
	"practise-code/utils"
)

func CheckUserName(name string) error  {
	var user types.User
	return global.SQLHARD.Db.Where("username = ?", name).First(&user).Error
}

func Register(user types.User)error{
	user.Password = utils.MD5(user.Password)
	return global.SQLHARD.Db.Create(&user).Error
}
