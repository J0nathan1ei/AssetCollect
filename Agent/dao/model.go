package dao

import "gorm.io/gorm"

type AppCode struct{
	gorm.Model

	AgentDomain		string	`json:"agent_domain"    gorm:"column:agent_domain;    type:varchar(100);    not null;	primaryKey;"`
	AgentAppcode	string	`json:"agent_appcode"   gorm:"column:appcode;         type:varchar(100);    not null;"`
}


func (i *AppCode) TableName() string{
	return "appcode"
}