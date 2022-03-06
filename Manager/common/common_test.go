package common

import (
	"fmt"
	"manager/config"
	"testing"
)

// 测试sha256哈希
func TestGetSha256Hash(t *testing.T) {
	ret := GetSha256Hash("123456")
	if ret != "8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92" {
		t.Error("sha256 value error!")
	}
}

// 测试密码盐值的生成
func TestGenSalt(t *testing.T) {
	salt1 := GenSalt()
	salt2 := GenSalt()
	if len(salt1) == 0 || len(salt2) == 0 {
		t.Error("generate salt failed")
	}
	if salt1 == salt2 {
		t.Error("random salts should not be the same")
	}
	fmt.Println(salt1)
	fmt.Println(salt2)
}

// 测试logger
func TestLogger(t *testing.T) {
	config.ManagerCfg.ZapCfg.Path = "../agent.log"
	config.ManagerCfg.ZapCfg.MaxSize = 10
	config.ManagerCfg.ZapCfg.MaxBackups = 5
	config.ManagerCfg.ZapCfg.MaxAge = 30
	config.ManagerCfg.ZapCfg.Compress = false
	InitLogger()
	Logger.Info("test info")
	Logger.Warn("test warn")
	Logger.Error("test error")
}

// 测试用户uuid
func TestGenUUID(t *testing.T) {
	uuid1 := GenUUID()
	uuid2 := GenUUID()
	if len(uuid1) == 0 || len(uuid2) == 0 || uuid1 == uuid2 {
		t.Error("gen uuid failed")
	}
	fmt.Println(uuid1)
	fmt.Println(uuid2)
}

// 测试校验器，校验非法字符
func TestValidator(t *testing.T) {
	type testStruct struct {
		Str1 string `validate:"required,ParamsValidation"`
	}
	var emptyDemo = testStruct{}
	var invalidDemo = testStruct{
		Str1: "？?{}",
	}
	var validDemo = testStruct{
		Str1: "admin",
	}

	InitValidator()
	// 空参数校验
	err := ParamsValidator.Struct(&emptyDemo)
	if err == nil {
		t.Error("validate empty params failed!")
	}
	fmt.Println("test empty param err: ", err)

	// 非法参数校验
	err = ParamsValidator.Struct(invalidDemo)
	if err == nil {
		t.Error("validate invalid params failed!")
	}
	fmt.Println("test invalid param err: ", err)

	// 合法参数校验
	err = ParamsValidator.Struct(validDemo)
	if err != nil {
		t.Error("validate valid params failed!")
	}
	fmt.Println("test valid param err: ", err)
}

func TestMapToString(t *testing.T) {
	var demo map[string]interface{} = make(map[string]interface{})
	demo["1"] = 1
	demo["2"] = 2
	ret, err := MapToString(demo)
	if err != nil {
		t.Error("map to string failed: ", err)
	}
	fmt.Println(ret)
}
