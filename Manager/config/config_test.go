package config

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var TestCfgPath string = "./config.yaml"

func TestInitConfig(t *testing.T) {

	Convey("错误读取配置", t, func() {
		err := InitConfig("xxx")
		So(err, ShouldNotBeNil)
		So(ManagerCfg, ShouldNotBeNil)
		So(ManagerCfg.PostgreCfg.User, ShouldEqual, "")
	})

	Convey("正常读取配置", t, func() {
		err := InitConfig(TestCfgPath)
		So(err, ShouldBeNil)
		So(ManagerCfg, ShouldNotBeNil)
		// 子项配置
		So(ManagerCfg.ZapCfg, ShouldNotBeNil)
		So(ManagerCfg.RedisCfg, ShouldNotBeNil)
		So(ManagerCfg.PostgreCfg, ShouldNotBeNil)
		So(ManagerCfg.ServerCfg, ShouldNotBeNil)
	})

}
