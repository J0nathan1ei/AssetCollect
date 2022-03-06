package controller

import (
	"encoding/json"
	"fmt"
	"manager/common"
	"manager/model"
	"net/http"
)

var loginResponse Response

func HandleLogin(w http.ResponseWriter, r *http.Request){
	var params ReqLogin
	var sessionId string
	var user model.User
	var retBytes []byte
	var exist bool

	// 解析登录请求
	reqStr := make([]byte, r.ContentLength)
	r.Body.Read(reqStr)
	err := json.Unmarshal(reqStr, &params)
	if err != nil{
		loginResponse.Success = false
		loginResponse.Msg = MsgErrParamsFormat + ": " + err.Error()
		goto END
	}

	// 参数检查
	if err = checkLoginParams(params); err != nil{
		loginResponse.Success = false
		loginResponse.Msg = err.Error()
		goto END
	}

	user.Username = params.UserName
	if exist, err = user.IsExist(); exist == false{
		loginResponse.Success = false
		loginResponse.Msg = MsgErrUserNotExist
		common.Logger.Error("Login error: ", err)
		goto END
	}

	// 开始登录
	sessionId, err = user.Login(params.Password)
	if err != nil{
		loginResponse.Success = false
		loginResponse.Msg = MsgErrWrongPwd
		goto END
	}

	loginResponse.Success = true
	loginResponse.Msg = MsgSuccessLogin + sessionId
END:
	retBytes, err = json.Marshal(loginResponse)
	if err != nil{
		common.Logger.Error("json marshaled failed")
		fmt.Fprintf(w, MsgErrServerInternal)
	}
	fmt.Fprintf(w, string(retBytes))
}