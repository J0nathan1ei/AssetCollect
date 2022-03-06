package config

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

var ManagerCfg *managerConfig = new(managerConfig)

type managerConfig struct {
	RedisCfg   redisConfig   `mapstructure:"redis"    validate:"required"`
	PostgreCfg postgreConfig `mapstructure:"postgre"  validate:"required"`
	ZapCfg     zapConfig     `mapstructure:"zap"      validate:"required"`
	ServerCfg  serverConfig  `mapstructure:"server"   validate:"required"`
}

type redisConfig struct {
	Addr     string `mapstructure:"addr"       validate:"required"`
	Password string `mapstructure:"password"   validate:"required"`
	DB       int    `mapstructure:"db"         validate:""`
}

type postgreConfig struct {
	Host     string `mapstructure:"host"          validate:"required"`
	Port     int    `mapstructure:"default_port"  validate:"required"`
	User     string `mapstructure:"user"          validate:"required"`
	Password string `mapstructure:"password"      validate:"required"`
	DB       string `mapstructure:"dbname"        validate:"required"`
	SSLMode  string `mapstructure:"ssl_mode"      validate:"required"`
	TimeZone string `mapstructure:"time_zone"     validate:"required"`
}

type zapConfig struct {
	Path       string `mapstructure:"file_name"     validate:"required"`
	MaxSize    int    `mapstructure:"max_size"      validate:"required"`
	MaxBackups int    `mapstructure:"max_backups"   validate:"required"`
	MaxAge     int    `mapstructure:"max_age"       validate:"required"`
	Compress   bool   `mapstructure:"compress"      validate:""`
}

type serverConfig struct {
	Mode string `mapstructure:"mode"       validate:"oneof=http https"`
}

// 初始化配置读取
var Path = "./config/config.yaml"

func InitConfig(path string) error{
	config := viper.New()
	config.SetConfigFile(path)
	if err := config.ReadInConfig(); err != nil {
		return errors.New("LOAD YAML FILE FAILED")
	}
	if err := config.Unmarshal(ManagerCfg); err != nil {
		return errors.New("UNMARSHAL YAML FAILED")
	}
	validate := validator.New()
	if err := validate.Struct(ManagerCfg); err != nil {
		return errors.New("CHECK YAML CONTENT FAILED")
	}
	return nil
}