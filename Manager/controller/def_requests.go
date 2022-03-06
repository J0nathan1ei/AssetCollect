package controller

type ReqLogin struct{
	UserName	string	`json:"username"  validate:"required,ParamsValidation,min=1,max=20"`
	Password	string	`json:"password"  validate:"required,ParamsValidation,min=1,max=50"`
}


type ReqRegister struct{
	UserName	string	`json:"username"  validate:"required,ParamsValidation,min=1,max=20"`
	Password	string	`json:"password"  validate:"required,ParamsValidation,min=1,max=50"`
	RePassword	string	`json:"re_password"  validate:"required,ParamsValidation,min=1,max=50"`
}