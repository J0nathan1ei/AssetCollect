package controller

// 失败信息
var (
	MsgErrParams 			string = "参数输入有误"
	MsgErrParamsFormat		string = "参数格式有误"
	MsgErrUserNotExist		string = "用户不存在"
	MsgErrWrongPwd			string = "密码输入错误"
	MsgErrDifferentPwd 		string = "两次密码输入不一致"
	MsgErrDaoInsertFailed	string = "数据插入失败"
	MsgErrServerInternal	string = "服务器内部错误，请稍后重试"

)

// 成功信息
var (
	MsgSuccessLogin			string = "登录成功"
	MsgSuccessRegister		string = "注册成功"
	MsgSuccessReport		string = "上报成功"
)