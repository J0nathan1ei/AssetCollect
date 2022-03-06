package collect

import (
	"agent/common"
	"errors"
	"fmt"
	"github.com/agiledragon/gomonkey"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAgentInfo_GetReportStr(t *testing.T) {
	var demo AgentInfo = AgentInfo{
		NewHostInfo(),
		NewNetInfo(),
		NewCpuInfo(),
		NewMemInfo(),
		NewDiskInfo(),
	}

	Convey("新建采集对象", t, func() {
		So(demo.Host, ShouldNotBeNil)
		So(demo.Net, ShouldNotBeNil)
		So(demo.CPU, ShouldNotBeNil)
		So(demo.Mem, ShouldNotBeNil)
		So(demo.Disk, ShouldNotBeNil)
	})
	Convey("生成上报结果", t, func() {
		ret := demo.GetReportStr()
		So(len(ret), ShouldNotBeZeroValue)
		fmt.Printf("%s\n", ret)
	})
	Convey("mock测试", t, func() {
		// 需要添加禁用内联函数选项 -gcflags all=-N -l

		// CPU信息
		patch1 := gomonkey.ApplyFunc(cpu.Info, func()([]cpu.InfoStat, error){
			return nil, errors.New("CPU info collect failed")
		})
		defer patch1.Reset()
		ret := demo.CPU.GetReportStr()
		So(len(ret), ShouldNotBeZeroValue)

		// Net信息
		patch2 := gomonkey.ApplyFunc(net.IOCounters, func(bool)([]net.IOCountersStat, error){
			return nil, errors.New("net info collect failed")
		})
		defer patch2.Reset()
		ret = demo.Net.GetReportStr()
		So(len(ret), ShouldNotBeZeroValue)

		// Host信息
		patch3 := gomonkey.ApplyFunc(host.Info, func()(*host.InfoStat, error){
			return nil, errors.New("host info collect failed")
		})
		defer patch3.Reset()
		ret = demo.Host.GetReportStr()
		So(len(ret), ShouldNotBeZeroValue)

		// Disk信息
		patch4 := gomonkey.ApplyFunc(disk.Usage, func(string)(*disk.UsageStat, error){
			return nil, errors.New("disk info collect failed")
		})
		defer patch4.Reset()
		ret = demo.Disk.GetReportStr()
		So(len(ret), ShouldNotBeZeroValue)

		// Mem信息
		patch5 := gomonkey.ApplyFunc(mem.VirtualMemory, func()(*mem.VirtualMemoryStat, error){
			return nil, errors.New("mem info collect failed")
		})
		defer patch5.Reset()
		ret = demo.Mem.GetReportStr()
		So(len(ret), ShouldNotBeZeroValue)
	})

}

func TestMain(m *testing.M){
	common.InitLogger()
	m.Run()
}
