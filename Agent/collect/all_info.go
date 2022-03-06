package collect

import "agent/common"

type AgentInfo struct {
	Host AgentAct `comment:"主机信息"`
	Net  AgentAct `comment:"网络信息"`
	CPU  AgentAct `comment:"CPU信息"`
	Mem  AgentAct `comment:"内存信息"`
	Disk AgentAct `comment:"硬盘信息"`
}

func NewAgentInfo() AgentAct {
	return &AgentInfo{
		NewHostInfo(),
		NewNetInfo(),
		NewCpuInfo(),
		NewMemInfo(),
		NewDiskInfo(),
	}
}

// 采集所有信息
func (i *AgentInfo) Collect() {
	i.Host.Collect()
	i.Net.Collect()
	i.CPU.Collect()
	i.Mem.Collect()
	i.Disk.Collect()
}

// 采集，然后生成上报字符串
func (i *AgentInfo) GetReportStr() string {
	i.Collect()
	return common.StructToJson(*i)
}
