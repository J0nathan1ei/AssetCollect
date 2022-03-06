package controller

import (
	"errors"
	"manager/common"
)

func commonFilter(in interface{})error{
	return  common.ParamsValidator.Struct(&in)
}


func checkLoginParams(in ReqLogin) error{
	// 公共校验
	if err := commonFilter(in); err!=nil{
		return errors.New(MsgErrParams)
	}

	return nil
}

func checkRegisterParams(in ReqRegister) error{
	// 公共校验
	if err := commonFilter(in); err!=nil{
		return errors.New(MsgErrParams)
	}

	// 两次密码一致性
	if in.Password != in.RePassword{
		return errors.New(MsgErrDifferentPwd)
	}

	return nil
}