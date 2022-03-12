package global

import (
	"go.uber.org/zap"
	"practise-code/config"
	rd "practise-code/core/redis"
	zp "practise-code/core/zap"
	"practise-code/model/sql"
)

var(
	CONFIG config.Config
	SQLHARD sql.SqlHard
	SLOG *zap.SugaredLogger
	REDIS rd.Redis
)

func InitGlobal()  {
	CONFIG.ReadConf() // 读取配置项
	SQLHARD.InitDB(CONFIG.Mysql)
	SQLHARD.AutoMigrate() // 自动迁移表格
	// 初始化zap logger -》SugaredLogger
	SLOG = zp.InitSugLogger(CONFIG.ZAP) // 配置日志打印文件
	REDIS = rd.InitRedis(CONFIG.Redis, SLOG)
}




