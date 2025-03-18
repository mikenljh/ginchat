package service

import (
	"ginchat/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUserList
// @Tags        用户模块
// @Summary     查看所有用户
// @Description 返回所有用户
// @Success      200 {string} json{"code","message"}
// @Router      /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()

	c.JSON(200, gin.H{
		"code":    0,
		"message": "success",
		"data":    data,
	})

}

// CreateUser
// @Tags        用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Summary     添加用户
// @Description 新增用户
// @Success      200 {string} json{"code","message"}
// @Router      /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")
	if password != repassword {
		c.JSON(200, gin.H{
			"message": "两次输入的密码不一致",
		})
		return
	}
	user.PassWord = password
	models.CreateUser(user)

	c.JSON(200, gin.H{
		"message": "新增用户成功",
	})

}

// DeleteUser
// @Tags        用户模块
// @param id query string false "id"
// @Summary     删除用户
// @Description 删除用户
// @Success      200 {string} json{"code","message"}
// @Router      /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"message": "删除用户成功",
	})

}

// UpdateUser
// @Tags        用户模块
// @param id formData string false "id"
// @param name formData string false "name"
// @param password formData string false "password"
// @Summary     修改用户
// @Description 修改用户
// @Success      200 {string} json{"code","message"}
// @Router      /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	models.UpdateUser(user)
	c.JSON(200, gin.H{
		"message": "修改用户成功",
	})

}
