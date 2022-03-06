package model

import (
	"gorm.io/gorm"
	"manager/common"
	"manager/dao"
	"time"
)

type User struct{
	gorm.Model

	Username	string		`gorm:"column:user_name;	type:varchar(100);	not null;"`
	PwdHash		string		`gorm:"column:pwd_hash;		type:text;			not null;"`
	Salt		string		`gorm:"column:salt;		type:varchar(100);	not null;"`
}

// 表名
func (this*User) TableName() string{
	return "users"
}

// 用户注册
func (this *User) Register(password string) error{
	this.Salt = common.GenSalt()
	this.PwdHash = common.GetSha256Hash(password + this.Salt)
	if err := dao.PostgreDB.Create(this).Error; err != nil{
		common.Logger.Error("Insert User Failed: ", err)
		return err
	}
	return nil
}

// 检查用户存在性
func (this *User) IsExist() (bool, error){
	var count int64
	err := dao.PostgreDB.Model(&User{}).Where("user_name = ?", this.Username).Count(&count).Error
	if err != nil{
		return false, err
	}
	return count == 1, nil
}

// 用户登录
var ExpireTime int = 600
func (this *User) Login(password string)(string, error){
	var user User
	var uuid string
	salt, err := this.getSalt()
	if err != nil{
		return "", err
	}
	pwdHash := common.GetSha256Hash(password + salt)
	err = dao.PostgreDB.Where("user_name = ? and pwd_hash = ?", this.Username, pwdHash).First(&user).Error
	if err != nil{
		return "", err
	}
	uuid = common.GenUUID()
	dao.RedisDB.Set(this.Username, uuid, time.Duration(ExpireTime)*time.Second)
	return uuid, nil
}

// 获取用户的盐值
func(this *User) getSalt()(string, error){
	var user User
	err := dao.PostgreDB.Where("user_name = ?", this.Username).First(&user).Error
	if err != nil{
		return "", err
	}
	return user.Salt, nil
}