package shortUrl

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuedun/ginode/db"
)

// GetLongByShort 根据短链获取长链
func GetLongByShort(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.(error).Error(),
			})
		}
	}()
	short := c.Query("short")
	shortService := NewService(db.Mysql)
	shortRecord, err := shortService.GetLongByShort(short)
	if err != nil {
		fmt.Println("err:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    shortRecord.OriginUrl,
		"message": "ok",
	})
}

// Long2Short 根据长链获取短链
func Long2Short(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.(error).Error(),
			})
		}
	}()
	long := c.Query("long")
	long2ShortRequest := Long2ShortRequest{OriginUrl: long}
	shortService := NewService(db.Mysql)
	short, err := shortService.Long2Short(&long2ShortRequest)
	if err != nil {
		fmt.Println("err:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    short,
		"message": "ok",
	})
}
