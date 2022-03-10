package global

import "practise-code/config"

var CONFIG config.Config

func InitGlobal()  {
	//CONFIG = config.Config{
	//	Port: "8080",
	//	JWT: config.JWT{
	//		SigningKey: "DSdf1s15afSDA5sd1fdf5sdf1f",
	//		ExpiresTime: 60*10, // 10分钟
	//		Issuer: "JP",
	//	},
	//}
	CONFIG.ReadConf()
}


