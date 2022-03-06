package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"manager/common"
	"manager/dao"
	"manager/model"
	"net/http"
)

// echo post请求的body
func EchoPost(w http.ResponseWriter, r *http.Request){
	requestBody, _ := ioutil.ReadAll(r.Body)

	// 返回值设置为json型
	w.Header().Set("content-type","text/json")

	fmt.Fprintf(w, string(requestBody))
}


func HandleAgentAllInfo(w http.ResponseWriter, r *http.Request){
	var data model.AgentInfo
	var response Response
	var err error
	var retBytes []byte
	requestBody, _ := ioutil.ReadAll(r.Body)

	// 返回值设置为json型
	w.Header().Set("content-type","text/json")

	if err := json.Unmarshal(requestBody, &data); err != nil{
		common.Logger.Error("Json Unmarshal Failed: ",err)
	}
	if err := dao.PostgreDB.Create(&data).Error; err != nil{
		common.Logger.Error(MsgErrDaoInsertFailed, err)
	}

	common.Logger.Info("Received Data: ", data)
	response.Success = true
	response.Msg = MsgSuccessReport

	retBytes, err = json.Marshal(response)
	if err != nil{
		common.Logger.Error("json marshaled failed")
		fmt.Fprintf(w, MsgErrServerInternal)
	}
	fmt.Fprintf(w, string(retBytes))
}


func HandleAgentNetInfo(w http.ResponseWriter, r *http.Request){
	EchoPost(w, r)
}

func HandleAgentCPUInfo(w http.ResponseWriter, r *http.Request){
	EchoPost(w, r)
}

func HandleAgentMemInfo(w http.ResponseWriter, r *http.Request){
	EchoPost(w, r)
}

func HandleAgentDiskInfo(w http.ResponseWriter, r *http.Request){
	EchoPost(w, r)
}

func HandleAgentHostInfo(w http.ResponseWriter, r *http.Request){
	EchoPost(w, r)
}
