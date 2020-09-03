package db

import (
	"gorm.io/gorm"
)

// Mysql 共享全局变量
var Mysql *gorm.DB
