package tests

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/yuedun/ginode/db"
	"github.com/yuedun/ginode/pkg/user"
)

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
