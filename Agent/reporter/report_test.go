package reporter

import (
	"agent/collect"
	"agent/common"
	"agent/config"
	"agent/dao"
	"fmt"
	"testing"
)

// 测试上报信息，需要开启mgr端

// 测试上报信息
func TestReport(t *testing.T) {
	testURL := "http://localhost:8001/user/v1/cpu-info"
	testCpuInfo := collect.CPUInfo{
		CPUModel: "Intel i9 10代",
		CPUCores: 64,
		CPUUsage: 5.1,
	}
	reportStr := common.StructToJson(testCpuInfo)
	report(reportStr, testURL)
}

// 测试
func TestAppCode(t *testing.T) {
	ret := getHeaderAppCode()
	if len(ret) == 0 {
		t.Error("get app code failed")
	}
	fmt.Println(ret)
}

// 测试上报所有信息
func TestReportAll(t *testing.T) {
	ReportAll()
}


func TestMain(m *testing.M){
	// 初始化
	var cfgPath string= "../config/config.yaml"
	common.InitLogger()
	config.InitConfig(cfgPath)
	dao.Init()

	m.Run()
}