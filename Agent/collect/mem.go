package collect

import (
	"agent/common"
	"fmt"
	"github.com/shirou/gopsutil/v3/mem"
)

type MemInfo struct {
	MemTotalSize string  `comment:"总内存大小"`
	MemSwapSize  string  `comment:"交换内存大小"`
	MemFreeSize  string  `comment:"剩余内存大小"`
	MemUsage     float64 `comment:"内存使用率"`
}

func NewMemInfo() AgentAct {
	return &MemInfo{}
}

func (i *MemInfo) Collect() {
	v, err := mem.VirtualMemory()
	if err != nil {
		common.Logger.Error("collect mem info failed: ", err)
		return
	}
	i.MemTotalSize = fmt.Sprintf("%dMB", v.Total/1024/1024)
	i.MemSwapSize = fmt.Sprintf("%dMB", v.SwapTotal/1024/1024)
	i.MemFreeSize = fmt.Sprintf("%dMB", v.Free/1024/1024)
	i.MemUsage = v.UsedPercent

	return
}

// 采集，然后生成上报字符串
func (i *MemInfo) GetReportStr() string {
	i.Collect()
	return common.StructToJson(*i)
}
