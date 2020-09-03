package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuedun/ginode/db"
	"github.com/yuedun/ginode/router"
	"github.com/yuedun/ginode/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	var err error
	c, _ := util.GetConf("conf.yaml")
	db.Mysql, err = gorm.Open(mysql.Open(fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Pwd, c.Host, c.Dbname)), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // SQL日志
	})
	if err != nil {
		panic(err)
	}
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
