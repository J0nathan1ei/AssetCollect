package main

import (
	"manager/common"
	"manager/config"
	"manager/dao"
	"manager/model"
	"manager/router"
)

func initTables() {
	// 初始化表，自动判断是否需要新建
	dao.PostgreDB.AutoMigrate(&model.AgentInfo{})
	dao.PostgreDB.AutoMigrate(&model.User{})
}

func initAll() {
	// 配置文件初始化
	config.InitConfig(config.Path)
	// 日志初始化
	common.InitLogger()
	// 校验器初始化
	common.InitValidator()
	// dao层初始化
	dao.InitAll()
	// 数据表初始化
	initTables()
}

func unInitAll() {
	dao.UnInit()
}

func main() {
	initAll()
	router.StartServer()
	unInitAll()
}
