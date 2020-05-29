package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/yuedun/ginode/util"
)

var Mysql *gorm.DB

func init() {
	var err error
	conf := util.Conf{}
	c := conf.GetConf()
	Mysql, err = gorm.Open("mysql", fmt.Sprintf("%v:%v@/%v?charset=utf8&parseTime=True&loc=Local", c.User, c.Pwd, c.Dbname))
	if err != nil {
		panic(err)
	}
	Mysql.LogMode(true)
	//Db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
	//defer Db.Close()
}
