package user

import (
	"practise-code/global"
	"practise-code/model/types"
)

func GetUserInfo(uid uint) ( userInfo types.User, err error) {
	err = global.SQLHARD.Db.Where("id = ?", uid).First(&userInfo).Error
	return userInfo, err
}
