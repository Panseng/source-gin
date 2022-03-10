package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

type Config struct {
	Name string `ini:"name"`
	Port string `ini:"port"`
	JWT
}
type JWT struct{
	SigningKey  string `ini:"signing_key"`    // jwt签名
	ExpiresTime int64  `ini:"expires_time"` // 过期时间，当前项目以秒为单位
	BufferTime  int64  `ini:"buffer_time"`    // 缓冲时间
	Issuer      string `ini:"issuer"`                  // 签发者
}

func (c *Config) ReadConf() {
	// 这里读取配置
	// 文件是相对 main.go 的位置
	cfg, err := ini.Load("config/cfg.ini")
	if err != nil{
		fmt.Printf("Fail to read file: %v", err)
	}
	c.Name = cfg.Section("").Key("name").String()
	c.Port =  cfg.Section("").Key("port").String()
	err = cfg.MapTo(c.JWT)
	if err != nil{
		fmt.Printf("Fail to map jwt: %v", err)
	}
}
