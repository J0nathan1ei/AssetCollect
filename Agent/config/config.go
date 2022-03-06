package config

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

var AgentCfg *agentConfig = new(agentConfig)

type agentConfig struct{
	Upload		upload			`mapstructure:"upload"      validate:"required"`
	Mysql   	mysqlConfig		`mapstructure:"mysql"       validate:"required"`
	ZapCfg		zapConfig		`mapstructure:"zap"         validate:"required"`
}

type upload struct{
	Url		url
}

type url struct{
	AllInfo	string		`mapstructure:"all_info"    validate:"required"`
	CpuInfo string		`mapstructure:"cpu_info"    validate:"required"`
	NetInfo string		`mapstructure:"net_info"    validate:"required"`
	HostInfo string		`mapstructure:"host_info"   validate:"required"`
	DiskInfo string		`mapstructure:"disk_info"   validate:"required"`
	MemInfo string		`mapstructure:"mem_info"    validate:"required"`
}

type mysqlConfig struct {
	User string			`mapstructure:"user"       validate:"required"`
	Pwd  string			`mapstructure:"password"   validate:"required"`
	Protocol string     `mapstructure:"protocol"   validate:"required"`
	Domain   string		`mapstructure:"domain"     validate:"required"`
	Port int			`mapstructure:"port"       validate:"required"`
	DBName string		`mapstructure:"db_name"    validate:"required"`
}


type zapConfig struct{
	Path		string		`mapstructure:"file_name"     validate:"required"`
	MaxSize		int			`mapstructure:"max_size"      validate:"required"`
	MaxBackups	int			`mapstructure:"max_backups"   validate:"required"`
	MaxAge		int			`mapstructure:"max_age"       validate:"required"`
	Compress	bool		`mapstructure:"compress"      validate:""`
}

// 初始化配置读取
var Path string = "./config/config.yaml"
func InitConfig(cfgPath string)error{
	fmt.Println(cfgPath)
	config := viper.New()
	config.SetConfigFile(cfgPath)
	if err := config.ReadInConfig(); err != nil{
		fmt.Errorf("Load yaml file failed: %s", err)
		return errors.New("LOAD YAML FAILED")
	}
	if err := config.Unmarshal(AgentCfg); err != nil{
		fmt.Errorf("Unmarshal yaml file failed: %s", err)
		return errors.New("UNMARSHAL YAML FAILED")
	}

	validate := validator.New()
	if err := validate.Struct(AgentCfg); err != nil {
		fmt.Println("Check yaml content failed: ", err)
		return errors.New("CHECK YAML CONTENT FAILED")
	}
	return nil
}