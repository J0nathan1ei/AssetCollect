package dao

import (
	"agent/config"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"ksuid"
)

var DB *gorm.DB

// 连接数据库并初始化
func Init()error{
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.AgentCfg.Mysql.User,
		config.AgentCfg.Mysql.Pwd,
		config.AgentCfg.Mysql.Protocol,
		config.AgentCfg.Mysql.Domain,
		config.AgentCfg.Mysql.Port,
		config.AgentCfg.Mysql.DBName)
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil{
		return errors.New("DB CONNECT FAILED")
	}
	DB = db
	// 建表
	createTables()

	return nil
}

// 创建Appcode表
func createTables(){
	DB.AutoMigrate(&AppCode{})
}


// 查询已有的Appcode
var LocalUser string = "localhost"
func GetAppcode(userDomain string)(string, error){
	var queryRet AppCode
	ret := DB.Where("agent_domain = ?", userDomain).First(&queryRet)
	if ret != nil{
		return queryRet.AgentAppcode, nil
	} else{
		return "", errors.New("USER APPCODE NOT EXISTS")
	}
}

// 首次生成AppCode
func CreateAppcode(userDomain string) string{
	firstAppcode := ksuid.New().String()
	DB.Create(&AppCode{
		gorm.Model{},
		userDomain,
		firstAppcode,
	})
	return firstAppcode
}
