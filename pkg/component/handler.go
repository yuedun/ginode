package component

import (
	"fmt"
	"github.com/yuedun/ginode/db"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//ComponentList列表
func ComponentList(c *gin.Context) {
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
	componentSearch := Component{
		Name:     name,
		Category: category,
		Status:   1,
	}
	wbService := NewService(db.Mysql)
	webs, err := wbService.GetComponentList(offset, limit, componentSearch)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    webs,
	})
}

//Create
func Create(c *gin.Context) {
	componentService := NewService(db.Mysql)
	wbObj := Component{}
	c.BindJSON(&wbObj)
	wbObj.CreatedAt = time.Now()
	wbObj.Status = 1
	err := componentService.CreateComponent(&wbObj)
	if err != nil {
		fmt.Println("err:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    wbObj,
		"message": "ok",
	})
}

//Update
func Update(c *gin.Context) {
	componentService := NewService(db.Mysql)
	component := Component{}
	c.BindJSON(&component)
	err := componentService.UpdateComponent(&component)
	if err != nil {
		fmt.Println("err:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data":    component,
			"message": "ok",
		})
	}
}

//Delete
func Delete(c *gin.Context) {
	componentId, _ := strconv.Atoi(c.Param("id"))
	componentService := NewService(db.Mysql)
	err := componentService.DeleteComponent(componentId)
	if err != nil {
		fmt.Println("err:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}
}