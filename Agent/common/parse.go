package common

import (
	"encoding/json"
	"time"
)


// 结构体信息转Json字符串
func StructToJson(in interface{})string{
	if in == nil{
		return ""
	}
	retBytes, err := json.Marshal(in)
	if err != nil{
		Logger.Error("Json Marshal Failed: ", err)
	}
	return string(retBytes)
}

// 获取当前时间
func GetNowTime()string{
	return time.Now().Format("2006-01-02 15:04:05")
}
