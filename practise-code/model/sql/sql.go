package sql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"practise-code/config"
	"practise-code/model/types"
)

// SqlH 持久化数据库连接
type SqlHard struct {
	Db *gorm.DB
}

func (s *SqlHard) InitDB(sqlCfg config.Mysql)  {
	dsn := sqlCfg.Username + ":"+ sqlCfg.Password+"@tcp("+sqlCfg.Path+":"+sqlCfg.Port+")/"+sqlCfg.Dbname+sqlCfg.Config
	if db, err := gorm.Open(mysql.New(mysql.Config{DSN: dsn})); err != nil{
		fmt.Println("================================================================")
		fmt.Println("mysql链接错误：" )
		fmt.Printf("mysql配置项：%+v\n", sqlCfg)
		fmt.Println("mysql配置文件位于config目录的config.ini" )
		fmt.Println(err.Error())
		fmt.Println("================================================================")
		return
	} else {
		s.Db = db
	}
}

func (s *SqlHard) AutoMigrate() {
	// 自动迁移表格
	s.Db.AutoMigrate(&types.User{})
}

