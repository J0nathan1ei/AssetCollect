package main

import (
	"agent/common"
	"agent/config"
	"agent/dao"
	"agent/reporter"
)


func initAll(){
	// 初始化日志
	common.InitLogger()
	// 初始化配置文件
	config.InitConfig(config.Path)
	// 初始化mysql数据库
	dao.Init()
}


func main() {
	initAll()
	reporter.ReportAll()
}
