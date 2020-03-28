package user

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type (
	/**
	  面向接口开发：
	  面向接口开发的好处是要对下面的函数进行测试时，不需要依赖一个全局的mysql连接，只需要调用NewService传一个mysql连接参数即可测试
	*/
	UserService interface {
		GetUserInfo(userObj User) (user User, err error)
		GetUserInfoBySQL() (user User, err error)
		CreateUser(user *User) (err error)
		UpdateUser(userID int, user *User) (err error)
		DeleteUser(userID int) (err error)
	}
)

type userService struct {
	mysql *gorm.DB
}

/*NewService 初始化结构体*/
func NewService(mysql *gorm.DB) UserService {
	return &userService{
		mysql: mysql,
	}
}

func (u *userService) GetUserInfo(userObj User) (user User, err error) {
	err = u.mysql.Where(userObj).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *userService) GetUserInfoBySQL() (user User, err error) {
	err = u.mysql.Raw("select * from user where id=?", user.Id).Scan(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *userService) CreateUser(user *User) (err error) {
	err = u.mysql.Create(user).Error
	fmt.Println(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) UpdateUser(userID int, user *User) (err error) {
	err = u.mysql.Model(user).Where("id = ?", userID).Update(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) DeleteUser(userID int) (err error) {
	u.mysql.Where("id = ?", userID).Delete(User{})
	if err != nil {
		return err
	}
	return nil
}

/**
面向接口开发的好处是，如果需要修改方法逻辑，可以在不修改原来逻辑的情况下新增接口实现，在调用的地方使用新的实现即可
*/
type newUserService struct {
	userService
	mysql *gorm.DB
}

/*NewUserService 初始化结构体*/
func NewUserService(mysql *gorm.DB) UserService {
	return &newUserService{
		mysql: mysql,
	}
}

func (u *newUserService) GetUserInfo(userObj User) (user User, err error) {
	err = u.mysql.First(&user).Error
	if err != nil {
		return user, err
	}
	fmt.Println("新方法")
	return user, nil
}
