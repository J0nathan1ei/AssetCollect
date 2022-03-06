package controller

// 标准请求响应
type Response struct{
	Success 	bool	`json:"success"`
	Msg			string	`json:"msg"`
}