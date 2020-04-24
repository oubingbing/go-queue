package util

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

/**
 * 加载redis配置信息
 */
func GetRedisConfig(redisClient *RedisClient){
	root, _ := os.Getwd()
	cfg,err := ini.Load(root+"/config.ini")
	if err != nil {
		Error(fmt.Sprintf("加载redis配置文件出错 -%v",err.Error()));
	}

	redisClient.Host     = cfg.Section("redis").Key("HOST").String()
	redisClient.Password = cfg.Section("redis").Key("PASSWORD").String()
	redisClient.DB       = cfg.Section("redis").Key("DB").MustInt()
}

/**
 * 加载mysql配置信息
 */
func GetMysqlConfig(db *Db)  {
	root, _ := os.Getwd()
	cfg,err := ini.Load(root+"/config.ini")
	if err != nil {
		Error(fmt.Sprintf("加载mysql配置文件出错 -%v",err.Error()));
	}

	db.Host     = cfg.Section("mysql").Key("DB_HOST").String()
	db.Port 	= cfg.Section("mysql").Key("DB_PORT").MustInt()
	db.Driver 	= cfg.Section("mysql").Key("DB_DRIVER").String()
	db.Username = cfg.Section("mysql").Key("DB_USERNAME").String()
	db.Password = cfg.Section("mysql").Key("DB_PASSWORD").String()
	db.Database = cfg.Section("mysql").Key("DB_DATABASE").String()
}

/**
 * 加载socket配置信息
 */
func GetSocketConfig() string {
	root, _ := os.Getwd()
	cfg,err := ini.Load(root+"/config.ini")
	if err != nil {
		Error(fmt.Sprintf("加载socket配置文件出错 -%v",err.Error()));
	}

	return cfg.Section("socket").Key("SOCKET_URL").String()
}

/**
 * 加载redis配置信息
 */
func GetCallBackConfig() string {
	root, _ := os.Getwd()
	cfg,err := ini.Load(root+"/config.ini")
	if err != nil {
		Error(fmt.Sprintf("加载call back配置文件出错 -%v",err.Error()));
	}

	return cfg.Section("redis").Key("REDIS_DB_CALLBACK_KEY").String()
}

/**
 * 加载log配置路径
 */
func GetLogConfig() string {
	root, _ := os.Getwd()
	cfg,err := ini.Load(root+"/config.ini")
	if err != nil {
		Error(fmt.Sprintf("加载logk配置文件出错 -%v",err.Error()));
	}

	return cfg.Section("log_path").Key("LOG_PATH").String()
}
