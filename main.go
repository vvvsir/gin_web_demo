package main

import (
	"fmt"
	"gin_web_demo/dao/mysql"
	"gin_web_demo/dao/redis"
	"gin_web_demo/logger"
	"gin_web_demo/pkg/snowflake"
	"gin_web_demo/router"
	"gin_web_demo/setting"
	"gin_web_demo/utils"
	"os"
)

// @title gin_web_demo
// @version 0.0.1
// @description  This is a api server.
// @BasePath /api/v1/
func main() {
	if len(os.Args) < 2 {
		fmt.Sprintln("need config file. eg: gin_web_demo comfig.yaml")
		return
	}
	// 加载配置
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("load config file failed, err:%v\n", err)
		return
	}
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("init logger config failed, err:%v\n", err)
		return
	}
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql config failed, err:%v\n", err)
		return
	}
	defer mysql.CLose()

	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("load redis config failed, err:%v\n", err)
		return
	}
	defer redis.Close()

	if err := snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	// 初始化gin框架内置的校验器使用的翻译器
	if err := utils.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed, err:%v\n", err)
		return
	}
	// 注册路由
	r := router.SetupRouter(setting.Conf.Mode)

	err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
