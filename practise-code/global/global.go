package global

import (
	"fmt"
	"practise-code/config"
	"practise-code/model/sql"
)

var(
	CONFIG config.Config
	SQLHARD sql.SqlHard
)

func InitGlobal()  {
	CONFIG.ReadConf() // 读取配置项
	fmt.Printf("配置：%+v\n", CONFIG)
	SQLHARD.InitDB(CONFIG.Mysql)
	SQLHARD.AutoMigrate() // 自动迁移表格
}


