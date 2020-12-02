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
	db.Mysql, err = gorm.Open(mysql.Open(fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Pwd, c.Host, c.Dbname)), &gorm.Config{
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
	userObj := user.User{Id: 1}
	user, err := userService.GetUserInfo(userObj)
	if err != nil {
		t.Error(err)
	}
	t.Log(user)
}

func TestCreateUser(t *testing.T) {
	userService := user.NewService(db.Mysql)
	newUser := new(user.User)
	newUser.Mobile = "17864345978"
	err := userService.CreateUser(newUser)
	if err != nil {
		t.Error(err)
	}
	t.Log(newUser)
}
