package model

import (
	"manager/common"
	"manager/config"
	"manager/dao"
	. "github.com/smartystreets/goconvey/convey"
	"gorm.io/gorm"
	"testing"
)


// 测试用户注册
func TestUser_Register(t *testing.T) {
	testUser := User{
		gorm.Model{},
		"jonathan",
		"123456",
		"",
	}
	Convey("用户正常注册", t, func() {
		err := testUser.Register("123456")
		So(err, ShouldBeNil)
	})
}


// 测试用户存在性
func TestUser_IsExist(t *testing.T) {
	testUser := User{
		gorm.Model{},
		"jonathan",
		"123456",
		"123456",
	}
	testInvalidUser := User{
		gorm.Model{},
		"hello",
		"456789",
		"456789",
	}
	Convey("已存在用户检查", t, func() {
		ret, err := testUser.IsExist()
		So(ret, ShouldBeTrue)
		So(err, ShouldBeNil)
	})
	Convey("不存在用户检查", t, func() {
		ret, err := testInvalidUser.IsExist()
		So(ret, ShouldBeFalse)
		So(err, ShouldBeNil)
	})

}

// 检查用户登录
func TestUser_Login(t *testing.T) {
	testUser := User{
		gorm.Model{},
		"jonathan",
		"123456",
		"",
	}
	Convey("用户正常登录", t, func() {
		ret, err :=testUser.Login("123456")
		So(len(ret), ShouldNotBeZeroValue)
		So(err, ShouldBeNil)
	})
	Convey("用户错误登录", t, func() {
		ret, err :=testUser.Login("xxxxxx")
		So(len(ret), ShouldBeZeroValue)
		So(err, ShouldNotBeNil)
	})

}

func TestMain(m *testing.M){
	_ = config.InitConfig("../config/config.yaml")
	dao.InitAll()
	common.InitLogger()
	m.Run()
}