package collect

import (
	"agent/common"
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"strconv"
	"time"
)

type CPUInfo struct {
	CPUModel string  `comment:"CPU型号"`
	CPUCores int     `comment:"CPU核数"`
	CPUUsage float64 `comment:"CPU使用率"`
}

func NewCpuInfo() AgentAct {
	return &CPUInfo{}
}

// CPU信息
func (i *CPUInfo) Collect() {
	pCpu, err := cpu.Info()
	if err != nil {
		common.Logger.Error("collect cpu info failed: ", err)
		return
	}
	i.CPUModel = pCpu[0].ModelName
	i.CPUCores = int(pCpu[0].Cores)
	cpuUsage, _ := cpu.Percent(time.Second, false)
	cpuUsageRet, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", cpuUsage[0]), 64)
	i.CPUUsage = cpuUsageRet

	return
}

// 采集，然后生成上报字符串
func (i *CPUInfo) GetReportStr() string {
	i.Collect()
	return common.StructToJson(*i)
}
