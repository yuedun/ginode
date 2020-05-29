package post

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/yuedun/ginode/db"

	"github.com/gin-gonic/gin"
)

//Index
func Index(c *gin.Context) {
	nameBody := map[string]string{}
	name := c.Request.Body
	nameByte, _ := ioutil.ReadAll(name)
	json.Unmarshal(nameByte, &nameBody)
	fmt.Println(nameBody)
	c.JSON(200, gin.H{
		"message": nameBody["name"],
	})
}

//GetPostInfo
func GetPostInfo(c *gin.Context) {
	userService := NewService(db.Mysql)
	user, err := userService.GetPostInfo()
	if err != nil {
		fmt.Println("err:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

//GetPostInfoBySql
func GetPostInfoBySql(c *gin.Context) {
	userService := NewService(db.Mysql)
	user, err := userService.GetPostInfoBySQL()
	if err != nil {
		fmt.Println("err:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

//CreatePost
func CreatePost(c *gin.Context) {
	userService := NewService(db.Mysql)
	user := Post{}
	fmt.Println(">>>", c.PostForm("mobile"))
	user.Mobile = c.PostForm("mobile")
	user.CreatedAt = time.Now()
	err := userService.CreatePost(&user)
	if err != nil {
		fmt.Println("err:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    user,
		"message": "ok",
	})
}

//UpdatePost
func UpdatePost(c *gin.Context) {
	userService := NewService(db.Mysql)
	user := Post{}
	userId, _ := strconv.Atoi(c.Param("id"))
	user.Addr = c.PostForm("addr")
	err := userService.UpdatePost(userId, &user)
	if err != nil {
		fmt.Println("err:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    user,
		"message": "ok",
	})
}

//DeletePost
func DeletePost(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	userService := NewService(db.Mysql)
	err := userService.DeletePost(userId)
	if err != nil {
		fmt.Println("err:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
