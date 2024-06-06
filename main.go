package main

import (
	"acana/dao/mysql"
	"acana/dao/redis"
	"acana/logger"
	"acana/setting"
	"fmt"
)

func main() {
	// step1: load config files
	if err := setting.Init(); err != nil {
		fmt.Printf("init setting failed, err: %v\n", err)
		return
	}
	// step2: record logs
	if err := logger.Init(setting.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err: %v\n", err)
		return
	}
	// step3: init mysql connect
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err: %v\n", err)
		return
	}
	// step4: init redis connect
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err: %v\n", err)
		return
	}
	// step5: register router
	// step6: start server

	fmt.Printf("End service\n")
}
