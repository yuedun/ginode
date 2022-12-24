package main

import (
	"flag"
	"net/http"
	"os/user"

	"github.com/gin-gonic/gin"
	"github.com/yuedun/ginode/db"
	"github.com/yuedun/ginode/pkg/component"
	"github.com/yuedun/ginode/pkg/post"
	"github.com/yuedun/ginode/pkg/shortUrl"
	"github.com/yuedun/ginode/pkg/website"
	"github.com/yuedun/ginode/router"
	"github.com/yuedun/ginode/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var config_path string

func init() {
	initFlag()
	flag.Parse()
	var err error
	c, _ := util.GetConf(config_path)
	db.Mysql, err = gorm.Open(mysql.Open(c.MysqlURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // SQL日志

	})
	if err != nil {
		panic(err)
	}
	db.Mysql.AutoMigrate(
		&component.Component{},
		&post.Post{},
		&user.User{},
		&shortUrl.ShortUrl{},
		&website.Website{},
	)
	util.InitLogger()
	defer util.SugarLogger.Sync()
}

func initFlag() {
	flag.StringVar(&config_path, "conf", "./conf.yaml", "config yaml path")
}

func main() {
	r := gin.Default()
	//r.Use(middleware.Logger())//全局中间件
	r.LoadHTMLGlob("templates/*") //加载模板
	r.GET("/", func(c *gin.Context) {
		util.SugarLogger.Debug(">>>>>>>")
		c.HTML(http.StatusOK, "index.tpl", gin.H{
			"title": "Hello Ginode!",
		})
	})

	router.Register(r)
	r.Run(":8909") // listen and serve on 0.0.0.0:8080
}
