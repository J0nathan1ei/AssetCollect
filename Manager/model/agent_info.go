package model

import "gorm.io/gorm"

// Agent信息
type AgentInfo struct{
	gorm.Model

	// 主机信息
	AgentId			string	`json:"AgentId"			gorm:"column:agent_id; 			type:varchar(100);	comment:agentID;	not null; 	primaryKey"`
	HostName		string	`json:"HostName"  		gorm:"column:host_name; 		type:text; 			comment:主机名;		not null;`
	HostPlatform	string	`json:"HostPlatform"  	gorm:"column:host_platform; 	type:text; 			comment:操作平台;		not null;`

	// 网络信息
	HostIP			string	`json:"HostIP"  		gorm:"column:host_ip;	 		type:text; 			comment:主机IP;		not null;`
	PacketsRcv		uint64	`json:"PacketsRcv"  	gorm:"column:packets_rcv; 		type:bigint; 		comment:网络收包数;	not null;`
	PacketsSend		uint64	`json:"PacketsSend"  	gorm:"column:packets_send;		type:bigint; 		comment:网络发包数;	not null;`

	// CPU信息
	CPUModel		string	`json:"CPUModel"  		gorm:"column:cpu_model; 		type:varchar(100);	comment:CPU型号;		not null;`
	CPUCores		int		`json:"CPUCores"  		gorm:"column:cpu_cores; 		type:smallint; 		comment:CPU核数;		not null;`
	CPUUsage		float64	`json:"CPUUsage"		gorm:"column:cpu_usage; 		type:decimal; 		comment:CPU使用率;	not null;`

	// 内存信息
	MemTotalSize	string	`json:"MemTotalSize"  	gorm:"column:mem_total_size;	type:varchar(30);	comment:总内存大小;	not null;`
	MemSwapSize		string	`json:"MemSwapSize"  	gorm:"column:mem_swap_size;		type:varchar(30);	comment:交换内存大小;	not null;`
	MemFreeSize		string	`json:"MemFreeSize"  	gorm:"column:mem_free_size;		type:varchar(30);	comment:剩余内存大小;	not null;`
	MemUsage		float64	`json:"MemUsage"  	   	gorm:"column:mem_usage; 		type:decimal(20,2); comment:内存使用率;	not null;`

	// 硬盘信息
	DiskTotalSize	string	`json:"DiskTotalSize"  	gorm:"column:disk_total_size;	type:varchar(30);	comment:总硬盘大小;	not null;`
	DiskUsage		float64	`json:"DiskUsage"  	   	gorm:"column:disk_usage; 		type:decimal(20,2); comment:硬盘使用率;	not null;`

}

func (AgentInfo) TableName() string {
	return "agent_info"
}
