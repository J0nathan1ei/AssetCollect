package collect

import (
	"agent/common"
	"github.com/shirou/gopsutil/v3/net"
)

type NetInfo struct {
	HostIP      string `comment:"主机IP"`
	PacketsRcv  uint64 `comment:"网络收包数"`
	PacketsSend uint64 `comment:"网络发包数"`
}

func NewNetInfo() AgentAct {
	return &NetInfo{}
}

// 网络信息
func (i *NetInfo) Collect() {
	nv, err := net.IOCounters(true)
	if err != nil {
		common.Logger.Error("collect net info failed: ", err)
		return
	}
	i.HostIP = getAllIp()
	i.PacketsRcv = nv[0].BytesRecv
	i.PacketsSend = nv[0].BytesSent

	return
}

//// 获取对外IP
//func getExternalIP() string {
//	conn, err := inet.Dial("udp", "8.8.8.8:80")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer conn.Close()
//
//	localAddr := conn.LocalAddr().(*inet.UDPAddr)
//	return localAddr.IP.String()
//}

// 获取所有网口IP
func getAllIp() string {
	var ret string
	netInfo, _ := net.Interfaces()
	for _, val := range netInfo {
		ret += val.Addrs[1].Addr + ","
	}
	return ret
}

// 采集，然后生成上报字符串
func (i *NetInfo) GetReportStr() string {
	i.Collect()
	return common.StructToJson(*i)
}
