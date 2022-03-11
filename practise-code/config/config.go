package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

type Config struct {
	Name string `ini:"name"`
	Port string `ini:"port"`
	DBType bool `ini:"db_type"` // 数据库类型 默认mysql，目前只支持mysql
	JWT
	Mysql
	MD5
}

type JWT struct{
	SigningKey  string `ini:"signing_key"`    // jwt签名
	ExpiresTime int64  `ini:"expires_time"` // 过期时间，当前项目以秒为单位
	BufferTime  int64  `ini:"buffer_time"`    // 缓冲时间
	Issuer      string `ini:"issuer"`                  // 签发者
}

type Mysql struct {
	Path         string ` ini:"path"`                             // 服务器地址
	Port         string `ini:"port"`                             // 端口
	Config       string `ini:"config"`                       // 高级配置
	Dbname       string `ini:"db_name"`                     // 数据库名
	Username     string `ini:"username"`                 // 数据库用户名
	Password     string `ini:"password"`                 // 数据库密码
	MaxIdleConns int    `ini:"max_idle_conns"` // 空闲中的最大连接数
	MaxOpenConns int    `ini:"max_open_conns"` // 打开到数据库的最大连接数
}

type MD5 struct {
	Salt string `ini:"salt"` // MD5 盐
}

func (c *Config) ReadConf() {
	// 这里读取配置
	// 文件是相对 main.go 的位置
	cfg, err := ini.Load("config/cfg.ini")
	if err != nil{
		fmt.Printf("Fail to read file: %v", err)
	}
	// 设置默认值
	// 这样设置也可以：c.JWT.ExpiresTime = 300
	c.JWT = JWT{
		ExpiresTime: 500,
	}
	// mysql 配置项
	c.Mysql = Mysql{
		Path: "localhost",
		Port: "3306",
		Config: "?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai",
	}
	c.MD5 = MD5{Salt: "1as2*DfS4^&5adSAda@DF%5"}

	err = cfg.MapTo(c)
	if err != nil{
		fmt.Printf("Fail to map jwt: %v", err)
	}
}
