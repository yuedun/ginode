package component

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/yuedun/ginode/db"

	"github.com/gin-gonic/gin"
)

type any = interface{}

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
	list, total, err := wbService.GetComponentList(offset, limit, componentSearch)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	data := map[string]any{
		"result": list,
		"count":  total,
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    data,
	})
}

//Create
func Create(c *gin.Context) {
	componentService := NewService(db.Mysql)
	wbObj := Component{}
	err := c.ShouldBind(&wbObj)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":    wbObj,
			"message": err.Error(),
		})
		return
	}
	wbObj.CreatedAt = time.Now()
	wbObj.Status = 1
	err = componentService.CreateComponent(&wbObj)
	if err != nil {
		fmt.Println("err:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    wbObj,
			"message": err.Error(),
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
	componentService := NewService(db.Mysql)
	component := Component{}
	c.ShouldBind(&component)
	err := componentService.UpdateComponent(&component)
	if err != nil {
		fmt.Println("err:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
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
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}
}
