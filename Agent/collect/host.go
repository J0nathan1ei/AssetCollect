package collect

import (
	"agent/common"
	"github.com/shirou/gopsutil/v3/host"
)

type HostInfo struct {
	AgentId      string `comment:"agentID"`
	HostName     string `comment:"主机名"`
	HostPlatform string `comment:"操作平台"`
}

func NewHostInfo() AgentAct {
	return &HostInfo{}
}

// 采集信息
func (i *HostInfo) Collect() {
	pHost, err := host.Info()
	if err != nil {
		common.Logger.Error("collect host info failed: ", err)
		return
	}
	i.AgentId = pHost.HostID
	i.HostName = pHost.Hostname
	i.HostPlatform = pHost.Platform

	return
}

// 采集，然后生成上报字符串
func (i *HostInfo) GetReportStr() string {
	i.Collect()
	return common.StructToJson(*i)
}
