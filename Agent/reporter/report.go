package reporter

import (
	"agent/collect"
	"agent/common"
	"agent/config"
	"agent/dao"
	"net/http"
	"strings"
)

// 上报全部信息
func ReportAll(){
	report(collect.NewAgentInfo().GetReportStr(), config.AgentCfg.Upload.Url.AllInfo)
	report(collect.NewHostInfo().GetReportStr(),  config.AgentCfg.Upload.Url.HostInfo)
	report(collect.NewCpuInfo().GetReportStr(),   config.AgentCfg.Upload.Url.CpuInfo)
	report(collect.NewNetInfo().GetReportStr(),   config.AgentCfg.Upload.Url.NetInfo)
	report(collect.NewMemInfo().GetReportStr(),   config.AgentCfg.Upload.Url.MemInfo)
	report(collect.NewDiskInfo().GetReportStr(),  config.AgentCfg.Upload.Url.DiskInfo)

}


// 上报信息
func report(info string, url string){
	common.Logger.Info("start uploading to url: ", url)
	req, err := http.NewRequest("POST", url, strings.NewReader(info))
	// 添加APPCODE字段
	req.Header.Add("Authorization", getHeaderAppCode())
	if err!=nil{
		common.Logger.Error("upload info failed: ", err)
	}
}

// 获取APPCODE
func getHeaderAppCode()string{
	appcode, err := dao.GetAppcode(dao.LocalUser)
	if err != nil{
		common.Logger.Info("First time Register, creating user appcode...")
		appcode = dao.CreateAppcode(dao.LocalUser)
	}
	return "APPCODE " + appcode
}