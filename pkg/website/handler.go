package website

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/yuedun/ginode/db"

	"github.com/gin-gonic/gin"
)

type any = interface{}

//WebsiteList列表
func WebsiteList(c *gin.Context) {
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		offset = 0
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}
	name := c.Query("name")
	category := c.Query("category")
	websiteSearch := Website{
		Name:     name,
		Category: category,
		Status:   1,
	}
	wbService := NewService(db.Mysql)
	list, total, err := wbService.GetWebsiteList(offset, limit, websiteSearch)
	data := map[string]any{
		"result": list,
		"count":  total,
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    data,
	})
}

//Create
func Create(c *gin.Context) {
	websiteService := NewService(db.Mysql)
	wbObj := Website{}
	c.ShouldBind(&wbObj)
	wbObj.CreatedAt = time.Now()
	wbObj.Status = 1
	err := websiteService.CreateWebsite(&wbObj)
	if err != nil {
		fmt.Println("err:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    wbObj,
			"message": "ok",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    wbObj,
		"message": "ok",
	})
}

//Update
func Update(c *gin.Context) {
	websiteService := NewService(db.Mysql)
	website := Website{}
	c.ShouldBind(&website)
	err := websiteService.UpdateWebsite(&website)
	if err != nil {
		fmt.Println("err:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data":    website,
			"message": "ok",
		})
	}
}

//Delete
func Delete(c *gin.Context) {
	websiteId, _ := strconv.Atoi(c.Param("id"))
	websiteService := NewService(db.Mysql)
	err := websiteService.DeleteWebsite(websiteId)
	if err != nil {
		fmt.Println("err:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}
}
