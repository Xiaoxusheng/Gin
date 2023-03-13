package server

import (
	"Gin/models"
	"Gin/utility"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// PingExample godoc
// @Summary 添加好友接口
// @Param account query string true "账号"
// @Param token header string true "token"
// @Schemes
// @Description 用户名 token 邮箱为必填
// @Tags 公共方法
// @Accept json
// @Produce json
// @Success 200 {string} {"code":200,"msg":"添加好友成功！"}
// @Router  /user/join    [get]
func JoinPrivate(c *gin.Context) {
	//好友的account
	account := c.Query("account")
	token := c.GetHeader("token")
	use, err := utility.ParseWithClaims(token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
		return
	}
	if account == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "参数错误！",
		})
		return
	}
	//验证添加的好友是否存在
	u, e := models.GetUserByaccount(account)
	if e != nil {
		log.Println(e)
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "账号不存在！",
		})
		return
	}
	//是否已经为好友
	other, err := models.GetOther(u.Indently, use.Indently)
	fmt.Println(other, err)
	if other && err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "已经互为好友！",
		})
		return
	}
	//房间号
	id := utility.GetRoomId()
	err = models.InsertUseridently(&models.User_room{use.Indently, id, time.Now().Unix(), time.Now().Unix(), "private"})
	if err != nil {
		log.Println("1", err)
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "系统错误！",
		})
		return
	}
	mc, err := models.GetUserByaccount(account)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "系统错误！",
		})
		return
	}
	err = models.InsertUseridently(&models.User_room{mc.Indently, id, time.Now().Unix(), time.Now().Unix(), "private"})
	if err != nil {
		log.Println("2", err)
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "系统错误！",
		})
		return
	}

	//
	f := models.CreateRoom(&models.Room_id{id, use.Indently, "private", time.Now().Unix(), use.Username, mc.Username})
	if f {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "添加好友成功！",
		})
	}
}

// 删除好友
// PingExample godoc
// @Summary  删除好友接口
// @Param account query string true "账号"
// @Param token header string true "token"
// @Schemes
// @Description 用户名 token 邮箱为必填
// @Tags 公共方法
// @Accept json
// @Produce json
// @Success 200 {string} {"code":200,"msg":"删除成功！"}
// @Router  /user/delete      [get]
func DelPrivate(c *gin.Context) {
	account := c.Query("account")
	token := c.GetHeader("token")
	use, err := utility.ParseWithClaims(token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
		return
	}
	if account == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "参数错误！",
		})
		return
	}
	//验证账号是否存在
	mc, e := models.GetUserByaccount(account)
	if e != nil {
		log.Println(e)
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "账号不存在！",
		})
		return
	}
	//是否已经为好友
	other, err := models.GetOther(mc.Indently, use.Indently)
	fmt.Println(other, err)
	if !(other && err == nil) {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "不是好友关系！",
		})
		return
	}
	//删除
	err = models.Del(use.Indently, mc.Indently)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "系统错误！",
		})
		return
	}
	room, err := models.GetRoom(use.Indently)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "系统错误！",
		})
		return
	}

	err = models.DelGroup(room)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "系统错误！",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功！",
	})
}
