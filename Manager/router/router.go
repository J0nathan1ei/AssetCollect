package router

import (
	"fmt"
	"manager/controller"
	"net/http"
)
// https选项
//var certPath = "./Server.crt"
//var keyPath  = "./Server.pem"

// 主页显示内容
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Welcome to Asset-Collect-Manager!")
}


func StartServer(){
	http.HandleFunc("/", 						HelloHandler)
	http.HandleFunc("/user/v1/login", 			controller.HandleLogin)
	http.HandleFunc("/user/v1/register", 		controller.HandleRegister)
	http.HandleFunc("/user/v1/echo_post", 		controller.EchoPost)
	http.HandleFunc("/user/v1/all-info", 		controller.HandleAgentAllInfo)
	http.HandleFunc("/user/v1/net-info", 		controller.HandleAgentNetInfo)
	http.HandleFunc("/user/v1/disk-info",		controller.HandleAgentDiskInfo)
	http.HandleFunc("/user/v1/mem-info", 		controller.HandleAgentMemInfo)
	http.HandleFunc("/user/v1/host-info", 		controller.HandleAgentHostInfo)
	http.HandleFunc("/user/v1/cpu-info", 		controller.HandleAgentCPUInfo)
	http.ListenAndServe(":8001", nil)
	//http.ListenAndServeTLS(":8001", certPath, keyPath, nil)
}
