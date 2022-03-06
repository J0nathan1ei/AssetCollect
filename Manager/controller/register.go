package controller

import (
	"encoding/json"
	"fmt"
	"manager/common"
	"manager/model"
	"net/http"
)

var registerResponse Response

func HandleRegister(w http.ResponseWriter, r *http.Request){
	var params ReqRegister
	var err error
	var user model.User
	var retBytes []byte

	// 解析地址添加请求
	reqStr := make([]byte, r.ContentLength)
	r.Body.Read(reqStr)
	err = json.Unmarshal(reqStr, &params)
	if err != nil{
		common.Logger.Error("Json Unmarshal Failed: ",err)
	}

	// 参数检查
	if err = checkRegisterParams(params); err != nil{
		registerResponse.Success = false
		registerResponse.Msg = err.Error()
		goto END
	}

	// 开始注册
	user.Username = params.UserName
	if err = user.Register(params.Password); err!=nil{
		registerResponse.Success = false
		registerResponse.Msg = MsgErrServerInternal
		goto END
	}

	registerResponse.Success = true
	registerResponse.Msg = MsgSuccessRegister
END:
	retBytes, err = json.Marshal(loginResponse)
	if err != nil{
		common.Logger.Error("json marshaled failed")
		fmt.Fprintf(w, MsgErrServerInternal)
	}
	fmt.Fprintf(w, string(retBytes))
}