package types

import "practise-code/model"

type User struct {
	model.DefaultField
	Username    string         `json:"userName" gorm:"not null;comment:用户登录名"`                                                        // 用户登录名
	Password    string         `json:"-"  gorm:"not null;comment:用户登录密码"`                                                             // 用户登录密码
	NickName    string         `json:"nickName" gorm:"column:nick_name;default:QMPlusUser;comment:用户昵称"`                                            // 用户昵称
	HeaderImg   string         `json:"headerImg" gorm:"column:header_img;default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	Phone       string         `json:"phone"  gorm:"comment:用户手机号"` // 用户角色ID
	Email       string         `json:"email"  gorm:"comment:用户邮箱"`  // 用户邮箱
}
