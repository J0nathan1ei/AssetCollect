package common

import (
	"agent/config"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type testStruct struct {
	Str1 string
	Num1 int
}

// 测试通用结构体转json字符串
func TestStructToJson(t *testing.T) {
	demo := new(testStruct)
	demo.Str1 = "hello"
	demo.Num1 = 123
	Convey("一般结构体转json", t, func() {
		ret := StructToJson(*demo)
		So(len(ret), ShouldNotBeZeroValue)
	})
}

// 测试当前时间获取
func TestGetNowTime(t *testing.T) {
	ret := GetNowTime()
	Convey("获取当前时间", t, func() {
		So(len(ret), ShouldNotBeZeroValue)
	})
	fmt.Println(ret)
}

// 测试Logger
func TestLogger(t *testing.T) {
	config.AgentCfg.ZapCfg.Path = "../agent.log"
	config.AgentCfg.ZapCfg.MaxSize = 10
	config.AgentCfg.ZapCfg.MaxBackups = 5
	config.AgentCfg.ZapCfg.MaxAge = 30
	config.AgentCfg.ZapCfg.Compress = false
	InitLogger()
	Logger.Info("test info")
	Logger.Warn("test warn")
	Logger.Error("test error")
}

func TestMain(m *testing.M){
	m.Run()
}