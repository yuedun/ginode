package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/yuedun/ginode/db"
	_ "github.com/yuedun/ginode/db"
	"github.com/yuedun/ginode/router"
	"github.com/yuedun/ginode/util"
)

func init() {
	var err error
	conf := util.Conf{}
	c, _ := conf.GetConf("conf.yaml")
	db.Mysql, err = gorm.Open("mysql", fmt.Sprintf("%v:%v@/%v?charset=utf8&parseTime=True&loc=Local", c.User, c.Pwd, c.Dbname))
	if err != nil {
		panic(err)
	}
	db.Mysql.LogMode(true)
	//Db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
	//defer Db.Close()
}

func main() {
	r := gin.Default()
	//r.Use(middleware.Logger())//全局中间件
	r.LoadHTMLGlob("templates/*") //加载模板
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tpl", gin.H{
			"title": "Hello World!",
		})
	})

	router.Register(r)
	r.Run(":8900") // listen and serve on 0.0.0.0:8080
}
