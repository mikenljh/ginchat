package main

import (
	"ginchat/router"
	"ginchat/utils"
)

func main() {
	//初始化配置文件以及数据库
	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()
	r := router.Router()
	r.Run(":9998")
}
