package dao

import (
	"fmt"
	"github.com/go-redis/redis"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"manager/common"
	"manager/config"
)

var (
	PostgreDB *gorm.DB
	RedisDB   *redis.Client
)

func InitAll(){
	initPostgre()
	initRedis()
}

// PostgreSql连接初始化
func initPostgre(){
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s timezone=%s",
		config.ManagerCfg.PostgreCfg.Host,
		config.ManagerCfg.PostgreCfg.Port,
		config.ManagerCfg.PostgreCfg.User,
		config.ManagerCfg.PostgreCfg.Password,
		config.ManagerCfg.PostgreCfg.DB,
		config.ManagerCfg.PostgreCfg.SSLMode,
		config.ManagerCfg.PostgreCfg.TimeZone)
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil{
		common.Logger.Error("PostgreSql Connect failed: ", err)
	}
	PostgreDB = db
}


// Redis连接初始化
func initRedis(){
	redisConn := redis.NewClient(&redis.Options{
		Addr:     config.ManagerCfg.RedisCfg.Addr,
		//Password: config.ManagerCfg.RedisCfg.Password,
		DB:       config.ManagerCfg.RedisCfg.DB,
	})
	if _, err := redisConn.Ping().Result(); err!=nil{
		common.Logger.Error("Redis Client build failed: ", err)
	}
	RedisDB = redisConn
}


// 程序退出时反初始化
func UnInit(){
	RedisDB.Close()
	RedisDB = nil
	PostgreDB = nil
}