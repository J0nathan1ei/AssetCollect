package dao

import (
	"agent/config"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

// 测试dao初始化
func TestInit(t *testing.T) {
	Convey("错误密码连接失败", t, func() {
		config.AgentCfg.Mysql.Pwd = "123"
		err := Init()
		So(DB, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})

	Convey("测试正常连接", t, func() {
		config.AgentCfg.Mysql.Pwd = "root"
		err := Init()
		So(DB, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

}

// 测试APPCODE
func TestCreateAppcode(t *testing.T) {
	Convey("测试APPCODE正常生成", t, func(){
		ret := CreateAppcode("8.8.8.8")
		So(len(ret), ShouldBeGreaterThan, 0)
	})
}

func TestMain(m *testing.M){
	var cfgPath string = "../config/config.yaml"
	config.InitConfig(cfgPath)
	m.Run()
}