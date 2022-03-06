package common

type ReqLogin struct{
	UserName	string	`json:"username"  validate:"required,ParamsValidation"`
	Password	string	`json:"password"  validate:"required,ParamsValidation"`
}


type ReqRegister struct{
	UserName	string	`json:"username"  validate:"required,ParamsValidation"`
	Password	string	`json:"password"  validate:"required,ParamsValidation"`
	RePassword	string	`json:"re_password"  validate:"required,ParamsValidation"`
}