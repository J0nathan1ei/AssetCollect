package collect

import (
	"agent/common"
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
)

type DiskInfo struct {
	DiskTotalSize string  `comment:"总硬盘大小"`
	DiskUsage     float64 `comment:"硬盘使用率"`
}

func NewDiskInfo() AgentAct {
	return &DiskInfo{}
}

func (i *DiskInfo) Collect() {
	d, err := disk.Usage("/")
	if err != nil {
		common.Logger.Error("collect disk info failed: ", err)
		return
	}
	i.DiskUsage = d.UsedPercent
	i.DiskTotalSize = fmt.Sprintf("%dGB", d.Total/1024/1024/1024)

	return
}

// 采集，然后生成上报字符串
func (i *DiskInfo) GetReportStr() string {
	i.Collect()
	return common.StructToJson(*i)
}
