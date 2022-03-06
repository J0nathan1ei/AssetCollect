package dao

import (
	. "github.com/smartystreets/goconvey/convey"
	"manager/config"
	"testing"
)

// 测试postgre初始化连接
func TestInitPostgre(t *testing.T) {
	Convey("Postgre连接", t, func() {
		So(PostgreDB, ShouldNotBeNil)
	})
}

// 测试redis初始化连接
func TestInitRedis(t *testing.T) {
	Convey("Redis连接", t, func() {
		So(RedisDB, ShouldNotBeNil)
	})

}

// 反初始化，关闭数据库连接
func TestUnInit(t *testing.T) {
	UnInit()
	Convey("关闭数据库连接", t, func() {
		So(PostgreDB, ShouldBeNil)
		So(RedisDB, ShouldBeNil)
	})
}

func TestMain(m *testing.M){
	var cfgPath string = "../config/config.yaml"
	config.InitConfig(cfgPath)
	InitAll()
	m.Run()
}