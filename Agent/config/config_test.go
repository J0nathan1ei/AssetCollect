package config

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var TestCfgPath string = "./config.yaml"

func TestInitConfig(t *testing.T) {

	Convey("错误配置读取", t, func() {
		err := InitConfig("xxx")
		So(err, ShouldNotBeNil)
		So(AgentCfg, ShouldNotBeNil)
	})

	Convey("正常配置读取", t, func() {
		err := InitConfig(TestCfgPath)
		So(err, ShouldBeNil)
		So(AgentCfg, ShouldNotBeNil)
	})

	Convey("检查配置内容", t, func(){
		// 上报url
		So(AgentCfg.Upload.Url.NetInfo, ShouldNotBeNil)
		So(AgentCfg.Upload.Url.HostInfo, ShouldNotBeNil)
		So(AgentCfg.Upload.Url.NetInfo, ShouldNotBeNil)
		So(AgentCfg.Upload.Url.MemInfo, ShouldNotBeNil)
		So(AgentCfg.Upload.Url.DiskInfo, ShouldNotBeNil)
		So(AgentCfg.Upload.Url.AllInfo, ShouldNotBeNil)

		//mysql 配置
		So(AgentCfg.Mysql.User, ShouldNotBeNil)
		So(AgentCfg.Mysql.Pwd, ShouldNotBeNil)
		So(AgentCfg.Mysql.Domain, ShouldNotBeNil)
		So(AgentCfg.Mysql.Port, ShouldNotBeNil)
		So(AgentCfg.Mysql.DBName, ShouldNotBeNil)
		So(AgentCfg.Mysql.Protocol, ShouldNotBeNil)

		// 日志配置
		So(AgentCfg.ZapCfg.MaxSize, ShouldNotBeNil)
		So(AgentCfg.ZapCfg.MaxAge, ShouldNotBeNil)
		So(AgentCfg.ZapCfg.MaxBackups, ShouldNotBeNil)
		So(AgentCfg.ZapCfg.Compress, ShouldNotBeNil)
		So(AgentCfg.ZapCfg.Path, ShouldNotBeNil)

	})
}


func TestMain(m *testing.M){
	m.Run()
}