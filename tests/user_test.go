package tests

import (
	"fmt"
	"testing"

	"github.com/yuedun/ginode/db"
	"github.com/yuedun/ginode/pkg/user"
	"github.com/yuedun/ginode/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//！！！！重要作用，用于初始化数据库
func TestMain(m *testing.M) {
	fmt.Println("begin")
	var err error
	c, _ := util.GetConf("../conf.yaml")
	db.Mysql, err = gorm.Open(mysql.Open(c.MysqlURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	m.Run()
	fmt.Println("end")
}
func TestGetUser(t *testing.T) {
	userService := user.NewService(db.Mysql)
	userObj := user.User{UserName: "月盾"}
	user, err := userService.GetUserInfo(userObj)
	if err != nil {
		t.Error(err)
	}
	t.Log(user)
}

func TestCreateUser(t *testing.T) {
	userService := user.NewService(db.Mysql)
	newUser := new(user.User)
	newUser.Mobile = "13333333333"
	newUser.UserName = "月盾"
	err := userService.CreateUser(newUser)
	if err != nil {
		t.Error(err)
	}
	t.Log(newUser)
}
