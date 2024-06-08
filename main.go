package main

import (
	"acana/dao/mysql"
	"acana/dao/redis"
	"acana/logger"
	"acana/pkg/snowflake"
	"acana/routes"
	"acana/setting"
	"fmt"

	"go.uber.org/zap"
)

func main() {
	// step1: load config files
	if err := setting.Init(); err != nil {
		fmt.Printf("init setting failed, err: %v\n", err)
		return
	}
	// step2: record logs
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err: %v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success")
	// step3: init mysql connect
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err: %v\n", err)
		return
	}
	defer mysql.Close()
	// step4: init redis connect
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err: %v\n", err)
		return
	}
	defer redis.Close()

	// init snowflake algos gen id
	if err := snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}

	// step5: register router
	r := routes.SetupRouter(setting.Conf.Mode)
	// step6: start server
	err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}

	fmt.Printf("End service\n")
}
